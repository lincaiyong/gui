package lsp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/lincaiyong/log"
	"github.com/mitchellh/mapstructure"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Message struct {
	JSONRPC string `json:"jsonrpc"`
	ID      any    `json:"id,omitempty"`
	Method  string `json:"method"`
	Params  any    `json:"params,omitempty"`
	Result  any    `json:"result,omitempty"`
	Error   *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    any    `json:"data,omitempty"`
	} `json:"error,omitempty"`
}

func TextDocumentParams(filePath string, lineIdx, charIdx int) any {
	type TextDocument struct {
		URI string `json:"uri"`
	}
	type Position struct {
		LineIdx int `json:"line"`
		CharIdx int `json:"character"`
	}
	params := struct {
		TextDocument TextDocument `json:"textDocument"`
		Position     Position     `json:"position"`
	}{
		TextDocument: TextDocument{
			URI: "file://" + filePath,
		},
		Position: Position{
			LineIdx: lineIdx,
			CharIdx: charIdx,
		},
	}
	return params
}

type Client struct {
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout io.ReadCloser
	msgID  int
	reader *bufio.Reader

	sourceRoot string
	files      map[string][]byte
}

func CreateClient() (*Client, error) {
	cmd := exec.Command("gopls")
	cmd.Env = append(os.Environ(), "GOPROXY=off", "GONOSUMDB=*", "GOPRIVATE=*")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err = cmd.Start(); err != nil {
		return nil, err
	}
	return &Client{
		cmd:    cmd,
		stdin:  stdin,
		stdout: stdout,
		msgID:  0,
		reader: bufio.NewReader(stdout),
	}, nil
}

func (c *Client) Close() {
	if c.stdin != nil {
		err := c.stdin.Close()
		if err != nil {
			log.WarnLog("fail to close stdin: %v", err)
		}
	}
	if c.stdout != nil {
		err := c.stdout.Close()
		if err != nil {
			log.WarnLog("fail to close stdout: %v", err)
		}
	}
	if c.cmd != nil && c.cmd.Process != nil {
		err := c.cmd.Process.Kill()
		if err != nil {
			log.WarnLog("fail to kill process: %v", err)
		}
	}
}

func (c *Client) nextID() int {
	c.msgID++
	return c.msgID
}

func (c *Client) sendRequest(method string, params any, withId bool) error {
	request := Message{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
	}
	if withId {
		request.ID = c.nextID()
	}
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return err
	}
	header := fmt.Sprintf("Content-Length: %d\r\n\r\n", len(requestBytes))
	if _, err = c.stdin.Write([]byte(header)); err != nil {
		return err
	}
	if _, err = c.stdin.Write(requestBytes); err != nil {
		return err
	}
	return nil
}

func (c *Client) safeReadLine() (line string, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.ErrorLog("safeReadLine panic: %v", r)
			return
		}
	}()
	return c.reader.ReadString('\n')
}

func (c *Client) readWithTimeout(timeout time.Duration) (string, error) {
	type result struct {
		line string
		err  error
	}

	resultChan := make(chan result, 1)

	go func() {
		line, err := c.safeReadLine()
		resultChan <- result{line: line, err: err}
	}()

	select {
	case res := <-resultChan:
		return res.line, res.err
	case <-time.After(timeout):
		return "", fmt.Errorf("read timeout after %v", timeout)
	}
}

func (c *Client) readMessage() (*Message, error) {
	headerLine, err := c.readWithTimeout(60 * time.Second)
	if err != nil {
		return nil, err
	}
	if !strings.HasPrefix(headerLine, "Content-Length:") {
		return nil, fmt.Errorf("expected Content-Length header, got: %s", headerLine)
	}
	lengthStr := strings.TrimSpace(strings.TrimPrefix(headerLine, "Content-Length:"))
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return nil, err
	}
	if _, err = c.safeReadLine(); err != nil {
		return nil, err
	}
	body := make([]byte, length)
	if _, err = io.ReadFull(c.reader, body); err != nil {
		return nil, err
	}
	var message Message
	err = json.Unmarshal(body, &message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (c *Client) readResponse() (*Message, error) {
	for {
		message, err := c.readMessage()
		if err != nil {
			return nil, err
		}
		// If this is a notification (no ID), ignore
		if message.ID == nil {
			log.DebugLog("gopls notification: %v\n", message.Params)
			continue
		}
		return message, nil
	}
}

func (c *Client) OpenProject(projectDir string) error {
	return c.initialize(projectDir)
}

func (c *Client) OpenFile(relPath string) error {
	if c.files == nil {
		c.files = make(map[string][]byte)
	}
	filePath := filepath.Join(c.sourceRoot, relPath)
	var b []byte
	var ok bool
	if b, ok = c.files[relPath]; !ok {
		var err error
		b, err = os.ReadFile(filePath)
		if err != nil {
			return err
		}
		c.files[relPath] = b
	}
	return c.didOpen(filePath, string(b))
}

type Position_ struct {
	Line      int
	Character int
}

type Target_ struct {
	URI   string
	Range struct {
		Start Position_
		End   Position_
	}
}

type Target struct {
	File       string `json:"file,omitempty"`
	LineIdx    int    `json:"line_idx,omitempty"`
	CharIdx    int    `json:"char_idx,omitempty"`
	EndLineIdx int    `json:"end_line_idx,omitempty"`
	EndCharIdx int    `json:"end_char_idx,omitempty"`
}

func (c *Client) QueryDefinition(relPath string, lineIdx, charIdx int) ([]*Target, error) {
	filePath := filepath.Join(c.sourceRoot, relPath)
	ret, err := c.getDefinition(filePath, lineIdx, charIdx)
	if err != nil {
		return nil, err
	}
	var result []*Target
	for _, item := range ret {
		var target_ Target_
		err = mapstructure.Decode(item, &target_)
		if err != nil {
			log.WarnLog("fail to decode target %v", err)
			continue
		}
		target := Target{
			File:       strings.TrimPrefix(target_.URI, "file://"),
			LineIdx:    target_.Range.Start.Line,
			CharIdx:    target_.Range.Start.Character,
			EndLineIdx: target_.Range.End.Line,
			EndCharIdx: target_.Range.End.Character,
		}
		result = append(result, &target)
	}
	return result, nil
}

func (c *Client) initialize(rootPath string) error {
	absPath, err := filepath.Abs(rootPath)
	if err != nil {
		return err
	}
	c.sourceRoot = absPath
	params := struct {
		ProcessID    int            `json:"processId"`
		RootURI      string         `json:"rootUri"`
		Capabilities map[string]any `json:"capabilities"`
	}{
		ProcessID: os.Getpid(),
		RootURI:   "file://" + absPath,
		Capabilities: map[string]any{
			"textDocument": map[string]any{
				"definition":     map[string]any{},
				"typeDefinition": map[string]any{},
				"implementation": map[string]any{},
			},
		},
	}
	err = c.sendRequest("initialize", params, true)
	if err != nil {
		return err
	}
	response, err := c.readResponse()
	if err != nil {
		return err
	}
	if response.Error != nil {
		return fmt.Errorf("initialize error: %s", response.Error.Message)
	}
	return c.sendRequest("initialized", map[string]any{}, false)
}

func (c *Client) didOpen(uri, content string) error {
	params := map[string]any{
		"textDocument": struct {
			URI        string `json:"uri"`
			LanguageID string `json:"languageId"`
			Version    int    `json:"version"`
			Text       string `json:"text"`
		}{
			URI:        "file://" + uri,
			LanguageID: "go",
			Version:    1,
			Text:       content,
		},
	}
	return c.sendRequest("textDocument/didOpen", params, false)
}

func (c *Client) getDefinition(filePath string, lineIdx, charIdx int) ([]any, error) {
	params := TextDocumentParams(filePath, lineIdx, charIdx)
	err := c.sendRequest("textDocument/definition", params, true)
	if err != nil {
		return nil, err
	}
	response, err := c.readResponse()
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, fmt.Errorf("definition error: %s", response.Error.Message)
	}
	if items, ok := response.Result.([]any); ok {
		return items, nil
	}
	if response.Result != nil {
		return []any{response.Result}, nil
	}
	return []any{}, nil
}

func (c *Client) getTypeDefinition(filePath string, lineNo, charNo int) ([]any, error) {
	params := TextDocumentParams(filePath, lineNo, charNo)
	err := c.sendRequest("textDocument/typeDefinition", params, true)
	if err != nil {
		return nil, err
	}
	response, err := c.readResponse()
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, fmt.Errorf("type definition error: %s", response.Error.Message)
	}
	if items, ok := response.Result.([]any); ok {
		return items, nil
	}
	if response.Result != nil {
		return []any{response.Result}, nil
	}
	return []any{}, nil
}

func (c *Client) getImplementation(filePath string, lineNo, charNo int) ([]any, error) {
	params := TextDocumentParams(filePath, lineNo, charNo)
	err := c.sendRequest("textDocument/implementation", params, true)
	if err != nil {
		return nil, err
	}
	response, err := c.readResponse()
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, fmt.Errorf("implementation error: %s", response.Error.Message)
	}
	if items, ok := response.Result.([]any); ok {
		return items, nil
	}
	if response.Result != nil {
		return []any{response.Result}, nil
	}
	return []any{}, nil
}
