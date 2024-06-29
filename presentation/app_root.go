package presentation

import (
	"embed"

	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/presentation/app"
	"github.com/samber/do"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

type AppRoot struct {
	todo   *app.TodoApp
	assets *embed.FS
}

func NewAppRoot(i *do.Injector, assets *embed.FS) *AppRoot {
	return &AppRoot{do.MustInvoke[*app.TodoApp](i), assets}
}

func (r *AppRoot) Run() {
	err := wails.Run(&options.App{
		Title:  "Wails Note App",
		Width:  1024,
		Height: 768,
		// For Linux: ウィンドウサイズを最適化
		// See https://github.com/wailsapp/wails/issues/2431
		MaxWidth:  3840,
		MaxHeight: 2160,
		AssetServer: &assetserver.Options{
			Assets: r.assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        nil,
		OnDomReady:       r.todo.OnDomReady,
		Bind:             []interface{}{r.todo},
		EnumBind: []interface{}{
			todo.StatusSeq,
			// dialog.AllDialogTypes,
			// dialog.AllDialogActionTypes,
			// dialog.AllDialogButtons,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
