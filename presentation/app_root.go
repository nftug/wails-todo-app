package presentation

import (
	"embed"

	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/presentation/app"
	"github.com/nftug/wails-todo-app/presentation/types/dialog"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"gorm.io/gorm"
)

type AppRoot struct {
	app   *app.App
	greet *app.GreetApp
	todo  *app.TodoApp
}

func NewAppRoot(app *app.App, greet *app.GreetApp, todo *app.TodoApp) *AppRoot {
	return &AppRoot{app, greet, todo}
}

func (r *AppRoot) Run(assets *embed.FS) {
	err := wails.Run(&options.App{
		Title:  "Wails Note App",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        r.app.Startup,
		Bind: []interface{}{
			r.app, r.greet, r.todo,
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

// Mock

type AppRootMock struct {
	DB   *gorm.DB
	Todo *app.TodoApp
}

func NewAppRootMock(db *gorm.DB, todo *app.TodoApp) *AppRootMock {
	return &AppRootMock{db, todo}
}
