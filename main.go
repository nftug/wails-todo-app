package main

import (
	"embed"
	"log"

	"github.com/nftug/wails-todo-app/domain/todo/enums"
	"github.com/nftug/wails-todo-app/infrastructure"
	"github.com/nftug/wails-todo-app/presentation"
	"github.com/nftug/wails-todo-app/presentation/app"
	"github.com/nftug/wails-todo-app/usecase"
	"github.com/samber/do"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	injector := do.New()
	{
		infrastructure.Inject(injector)
		usecase.Inject(injector)
		presentation.Inject(injector)
	}

	todoApp := do.MustInvoke[*app.TodoApp](injector)

	if err := wails.Run(&options.App{
		Title:  "Wails Note App",
		Width:  1024,
		Height: 768,
		// For Linux: ウィンドウサイズを最適化
		// See https://github.com/wailsapp/wails/issues/2431
		MaxWidth:  3840,
		MaxHeight: 2160,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        nil,
		OnDomReady:       todoApp.OnDomReady,
		Bind:             []interface{}{todoApp},
		EnumBind: []interface{}{
			enums.StatusSeq,
			// dialog.AllDialogTypes,
			// dialog.AllDialogActionTypes,
			// dialog.AllDialogButtons,
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				FullSizeContent: true,
			},
		},
	}); err != nil {
		log.Fatal("Error:", err.Error())
	}
}
