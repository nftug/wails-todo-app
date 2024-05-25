package main

import (
	"embed"

	"github.com/nftug/wails-todo-app/presentation"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	presentation.CreateAppRoot().Run(&assets)
}
