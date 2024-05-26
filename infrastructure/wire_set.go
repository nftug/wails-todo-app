package infrastructure

import (
	"github.com/google/wire"
	"github.com/nftug/wails-todo-app/infrastructure/config"
	"github.com/nftug/wails-todo-app/infrastructure/persistence"
	"github.com/nftug/wails-todo-app/infrastructure/todo"
)

var commonSet = wire.NewSet(
	todo.NewTodoRepository,
	todo.NewTodoQueryService,
)

var Set = wire.NewSet(
	config.NewLocalPathService,
	persistence.NewDB,
	commonSet,
)

var MockSet = wire.NewSet(
	persistence.NewDBMock,
	commonSet,
)
