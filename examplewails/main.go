package main

import (
	"context"
	"embed"
	"github.com/lincaiyong/log"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed build/appicon.png
var icon []byte

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	log.EnableDebugLog()
	_ = log.SetLogPath("/tmp/codeedge.log")
	err := wails.Run(&options.App{
		Title:  "CodeEdge",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarDefault(),
			Appearance:           mac.DefaultAppearance,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "CodeEdge",
				Message: "©2025 lincaiyong@codeedge.cc",
				Icon:    icon,
			},
		},
		OnStartup: func(ctx context.Context) {
			app.ctx = ctx
		},
		Bind: []any{&app},
	})
	if err != nil {
		log.ErrorLog("fail to start: %v", err)
	}
}
