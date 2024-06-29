package main

import (
	"embed"

	"github.com/nftug/wails-todo-app/infrastructure"
	"github.com/nftug/wails-todo-app/presentation"
	"github.com/nftug/wails-todo-app/usecase"
	"github.com/samber/do"
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
	app := presentation.NewAppRoot(injector, &assets)
	app.Run()
}
