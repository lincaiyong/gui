package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lincaiyong/log"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"wails/lsp"
)

var app App

type App struct {
	ctx       context.Context
	lspClient *lsp.Client
}

func (a *App) Log(v any) {
	log.DebugLog("%v", v)
}

func (a *App) LspOpenProject(dir string) {
	log.DebugLog("lsp open project: %s", dir)
	err := a.lspClient.OpenProject(dir)
	if err != nil {
		log.ErrorLog("fail to open project: %v", err)
	}
}

func (a *App) LspOpenFile(file string) {
	log.DebugLog("lsp open file: %s", file)
	err := a.lspClient.OpenFile(file)
	if err != nil {
		log.ErrorLog("fail to open file: %v", err)
	}
}

func (a *App) LspQueryDefinition(file string, lineIdx, charIdx int) string {
	log.DebugLog("lsp query definition: %s#%d:%d", file, lineIdx+1, charIdx+1)
	targets, err := a.lspClient.QueryDefinition(file, lineIdx+1, charIdx)
	if err != nil {
		log.ErrorLog("fail to open file: %v", err)
	}
	b, _ := json.MarshalIndent(targets, "", "  ")
	return string(b)
}

func (a *App) ReadFile(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		log.ErrorLog("fail to read file %s: %v", err)
		return fmt.Sprintf("fail to read file %s: %v", path, err)
	}
	return string(b)
}

func (a *App) SelectFolder() string {
	log.DebugLog("select folder...")
	opts := runtime.OpenDialogOptions{
		Title: "选择项目",
	}
	folder, err := runtime.OpenDirectoryDialog(a.ctx, opts)
	if err != nil {
		log.ErrorLog("fail to OpenDirectoryDialog: %v", err)
		return ""
	}
	log.DebugLog("folder: %s", folder)
	var files []string
	err = filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && strings.HasPrefix(d.Name(), ".") {
			return fs.SkipDir
		}
		if !d.IsDir() {
			relPath, _ := filepath.Rel(folder, path)
			files = append(files, relPath)
		}
		return nil
	})
	data := map[string]any{
		"folder": folder,
		"files":  files,
	}
	ret, _ := json.Marshal(data)
	log.DebugLog("result: %s", string(ret))
	return string(ret)
}
