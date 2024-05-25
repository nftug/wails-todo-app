package presentation

import (
	"github.com/google/wire"
	"github.com/nftug/wails-todo-app/presentation/app"
)

var Set = wire.NewSet(
	app.NewApp,
	app.NewGreetApp,
	app.NewTodoApp,
	NewAppRoot,
)

var MockSet = wire.NewSet(
	app.NewTodoApp,
	NewAppRootMock,
)
