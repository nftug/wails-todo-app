package presentation

import (
	"embed"

	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/presentation/app"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"gorm.io/gorm"
)

type AppRoot struct {
	todo *app.TodoApp
}

func NewAppRoot(todo *app.TodoApp) *AppRoot {
	return &AppRoot{todo}
}

func (r *AppRoot) Run(assets *embed.FS) {
	err := wails.Run(&options.App{
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

// Mock

type AppRootMock struct {
	DB   *gorm.DB
	Todo *app.TodoApp
}

func NewAppRootMock(db *gorm.DB, todo *app.TodoApp) *AppRootMock {
	return &AppRootMock{db, todo}
}
