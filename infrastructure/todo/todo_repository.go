package todo

import (
	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/infrastructure/persistence/model"
	"gorm.io/gorm"
)

type TodoRepository struct {
	*model.Repository[*todo.Todo, *TodoTable]
}

func NewTodoRepository(db *gorm.DB) todo.TodoRepository {
	return &TodoRepository{model.NewRepository[*todo.Todo, *TodoTable](db)}
}
