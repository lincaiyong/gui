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
)

var app App

type App struct {
	ctx context.Context
}

//func logOutput(ctx context.Context, msg string) {
//	runtime.EventsEmit(ctx, "output", msg)
//}

func (a *App) Log(v any) {
	log.DebugLog("%v", v)
}

func (a *App) Reload() {
	log.DebugLog("reload")
	runtime.EventsEmit(a.ctx, "wails:reload:frontend")
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
