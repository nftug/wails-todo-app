package todo

import (
	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/infrastructure/persistence"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
	*persistence.Repository[*todo.Todo, *TodoDBSchema]
}

func NewTodoRepository(i *do.Injector) (todo.TodoRepository, error) {
	db := do.MustInvoke[*gorm.DB](i)
	return &todoRepository{db, persistence.NewRepository[*todo.Todo, *TodoDBSchema](i)}, nil
}
