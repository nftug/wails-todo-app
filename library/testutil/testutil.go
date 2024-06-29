package testutil

import (
	"testing"

	"github.com/nftug/wails-todo-app/infrastructure"
	"github.com/nftug/wails-todo-app/usecase"
	"github.com/samber/do"
)

func GetInjector(t *testing.T) *do.Injector {
	t.Helper()

	injector := do.New()
	{
		infrastructure.InjectForTest(t, injector)
		usecase.Inject(injector)
	}
	return injector
}
