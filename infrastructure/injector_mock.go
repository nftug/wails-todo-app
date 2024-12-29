package infrastructure

import (
	"testing"

	"github.com/nftug/wails-todo-app/infrastructure/common/config"
	"github.com/samber/do"
)

func InjectForTest(t *testing.T, i *do.Injector) {
	t.Helper()

	injectCore(i)
	do.Provide(i, func(i *do.Injector) (config.ConfigPathService, error) {
		return config.NewConfigPathMockService(t)
	})
}
