package infrastructure

import (
	"github.com/google/wire"
	"github.com/nftug/wails-todo-app/infrastructure/config"
	"github.com/nftug/wails-todo-app/infrastructure/persistence"
	"github.com/nftug/wails-todo-app/infrastructure/todo"
)

var Set = wire.NewSet(
	config.NewLocalPathService,
	persistence.NewDB,
	todo.NewTodoRepository,
	todo.NewTodoQueryService,
)
