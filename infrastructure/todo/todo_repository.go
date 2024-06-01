package todo

import (
	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/infrastructure/persistence/model"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
	*model.Repository[*todo.Todo, *TodoDBSchema]
}

func NewTodoRepository(db *gorm.DB) todo.TodoRepository {
	return &TodoRepository{db, model.NewRepository[*todo.Todo, *TodoDBSchema](db)}
}
