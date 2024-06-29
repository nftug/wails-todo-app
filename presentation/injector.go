package presentation

import (
	"github.com/nftug/wails-todo-app/presentation/app"
	"github.com/samber/do"
)

func Inject(i *do.Injector) {
	do.Provide(i, app.NewTodoApp)
}
