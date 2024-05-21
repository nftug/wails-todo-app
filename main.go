package main

import (
	"embed"

	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/presentation/app"
	"github.com/nftug/wails-todo-app/presentation/types/dialog"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	a := app.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Wails Note App",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        a.Startup,
		Bind: []interface{}{
			a,
		},
		EnumBind: []interface{}{
			todo.StatusSeq,
			dialog.AllDialogTypes,
			dialog.AllDialogActionTypes,
			dialog.AllDialogButtons,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
