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

func (a *App) QueryDefinition(file string, lineIdx, charIdx int) string {
	log.DebugLog("lsp query definition: %s#%d:%d", file, lineIdx+1, charIdx+1)
	targets, err := a.lspClient.QueryDefinition(file, lineIdx, charIdx)
	if err != nil {
		log.ErrorLog("fail to open file: %v", err)
	}
	b, _ := json.Marshal(targets)
	log.InfoLog("targets: %s", string(b))
	return string(b)
}

func (a *App) OpenFile(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		log.ErrorLog("fail to read file %s: %v", err)
		return fmt.Sprintf("fail to read file %s: %v", path, err)
	}
	err = a.lspClient.OpenFile(path)
	if err != nil {
		log.ErrorLog("fail to open file: %v", err)
	}
	return string(b)
}

func (a *App) OpenDirectory() string {
	log.DebugLog("open directory...")
	opts := runtime.OpenDialogOptions{
		Title: "打开目录",
	}
	folder, err := runtime.OpenDirectoryDialog(a.ctx, opts)
	if err != nil {
		log.ErrorLog("fail to open: %v", err)
		return ""
	}
	return folder
}

func (a *App) OpenProject(folder string) string {
	log.DebugLog("open project: %s", folder)
	if app.lspClient != nil {
		app.lspClient.Close()
	}
	var err error
	app.lspClient, err = lsp.CreateClient()
	if err != nil {
		log.ErrorLog("fail to create lsp client: %v", err)
	}
	err = a.lspClient.OpenProject(folder)
	if err != nil {
		log.ErrorLog("fail to open project: %v", err)
	}
	var files []string
	err = filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && strings.HasPrefix(d.Name(), ".") {
			return fs.SkipDir
		}
		if !d.IsDir() && !strings.HasPrefix(d.Name(), ".") {
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
