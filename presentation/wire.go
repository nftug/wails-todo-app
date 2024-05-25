//go:build wireinject
// +build wireinject

package presentation

import (
	"github.com/google/wire"
	"github.com/nftug/wails-todo-app/infrastructure"
	"github.com/nftug/wails-todo-app/usecase"
)

func CreateAppRoot() *AppRoot {
	wire.Build(
		infrastructure.Set,
		usecase.Set,
		Set,
	)
	return &AppRoot{}
}

func CreateAppRootMock() *AppRootMock {
	wire.Build(
		infrastructure.MockSet,
		usecase.Set,
		MockSet,
	)
	return &AppRootMock{}
}
