package infrastructure

import (
	"testing"

	"github.com/nftug/wails-todo-app/infrastructure/middleware"
	"github.com/samber/do"
)

func InjectForTest(t *testing.T, i *do.Injector) {
	t.Helper()
	injectCore(i)
	do.Provide(i, middleware.NewDBMock)
}
