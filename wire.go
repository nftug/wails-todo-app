//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/nftug/wails-todo-app/infrastructure"
	"github.com/nftug/wails-todo-app/presentation"
	"github.com/nftug/wails-todo-app/usecase"
)

func createAppRoot() *presentation.AppRoot {
	wire.Build(
		infrastructure.Set,
		usecase.Set,
		presentation.Set,
	)
	return &presentation.AppRoot{}
}
