//go:build wireinject
// +build wireinject

package usecase

import (
	"github.com/google/wire"
	"github.com/nftug/wails-todo-app/infrastructure"
)

func InitUseCaseAdapterMock() *UseCaseAdapterMock {
	wire.Build(infrastructure.MockSet, Set)
	return &UseCaseAdapterMock{}
}
