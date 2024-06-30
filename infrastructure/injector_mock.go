package infrastructure

import (
	"testing"

	"github.com/nftug/wails-todo-app/infrastructure/middleware"
	"github.com/samber/do"
	"gorm.io/driver/sqlite"
)

func InjectForTest(t *testing.T, i *do.Injector) {
	t.Helper()

	injectCore(i)
	do.Provide(i, func(i *do.Injector) (middleware.ConfigPathService, error) {
		return middleware.NewConfigPathMockService(t)
	})
	do.ProvideValue(i, sqlite.Open(":memory:"))
}
