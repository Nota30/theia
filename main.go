package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	timer := NewTimer()

	err := wails.Run(&options.App{
		Title:  "Theia",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		StartHidden: true,
		AlwaysOnTop: true,
		Fullscreen: true,
		BackgroundColour: &options.RGBA{R: 35, G: 43, B: 43, A: 1},
		OnStartup:        timer.timerStartup,
		Bind: []interface{}{
			timer,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
