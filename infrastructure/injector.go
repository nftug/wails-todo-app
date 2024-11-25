package infrastructure

import (
	"github.com/nftug/wails-todo-app/infrastructure/common"
	"github.com/nftug/wails-todo-app/infrastructure/todo"
	"github.com/samber/do"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Inject(i *do.Injector) {
	injectCore(i)
	do.Provide(i, common.NewConfigPathService)
	do.Provide(i, func(i *do.Injector) (gorm.Dialector, error) {
		lp := do.MustInvoke[common.ConfigPathService](i)
		return sqlite.Open(lp.GetJoinedPath("todo.db")), nil
	})
	do.Provide(i, todo.NewTodoNotificationSender)
}

func injectCore(i *do.Injector) {
	do.Provide(i, common.NewDB)
	do.Provide(i, todo.NewTodoRepository)
	do.Provide(i, todo.NewTodoQueryService)
}
