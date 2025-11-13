package lsp

import (
	"fmt"
	"github.com/lincaiyong/log"
	"testing"
)

func testClient(projectDir, filePath string) error {
	c, err := CreateClient()
	if err != nil {
		return fmt.Errorf("fail to create lsp client: %v", err)
	}
	defer c.Close()
	err = c.OpenProject(projectDir)
	if err != nil {
		return fmt.Errorf("fail to initialize lsp client: %v", err)
	}
	err = c.OpenFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to send didopen notification: %s, %v", filePath, err)
	}
	targets, err := c.QueryDefinition(filePath, 16, 12)
	if err != nil {
		return fmt.Errorf("fail to query: %v", err)
	}
	for _, target := range targets {
		log.InfoLog("%s:%d", target.File, target.LineIdx+1)
	}
	return nil
}

func TestFoo(t *testing.T) {
	err := testClient("/Users/bytedance/Code/lincaiyong/gui/example", "example.go")
	if err != nil {
		t.Fatalf("fail to create lsp client: %v", err)
	}
}
