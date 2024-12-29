package infrastructure

import (
	"github.com/nftug/wails-todo-app/infrastructure/common/config"
	"github.com/nftug/wails-todo-app/infrastructure/common/db"
	"github.com/nftug/wails-todo-app/infrastructure/todo"
	"github.com/samber/do"
)

func Inject(i *do.Injector) {
	injectCore(i)
	do.Provide(i, config.NewConfigPathService)
	do.Provide(i, todo.NewTodoNotificationSender)
}

func injectCore(i *do.Injector) {
	do.Provide(i, db.NewBBoltDB)
	do.Provide(i, todo.NewTodoRepository)
	do.Provide(i, todo.NewTodoQueryService)
}
