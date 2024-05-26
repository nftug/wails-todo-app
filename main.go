package main

import (
	"embed"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	InitAppRoot().Run(&assets)
}
