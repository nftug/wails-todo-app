package infrastructure

import (
	"github.com/nftug/wails-todo-app/infrastructure/middleware"
	"github.com/nftug/wails-todo-app/infrastructure/todo"
	"github.com/samber/do"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Inject(i *do.Injector) {
	injectCore(i)
	do.Provide(i, middleware.NewConfigPathService)
	do.Provide(i, func(i *do.Injector) (gorm.Dialector, error) {
		lp := do.MustInvoke[middleware.ConfigPathService](i)
		return sqlite.Open(lp.GetJoinedPath("todo.db")), nil
	})
}

func injectCore(i *do.Injector) {
	do.Provide(i, middleware.NewDB)
	do.Provide(i, todo.NewTodoRepository)
	do.Provide(i, todo.NewTodoQueryService)
}
