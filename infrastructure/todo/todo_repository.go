package todo

import (
	"context"
	"log"

	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/infrastructure/persistence/model"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
	*model.Repository[*todo.Todo, TodoTable]
}

func NewTodoRepository(db *gorm.DB) todo.TodoRepository {
	return &TodoRepository{db, model.NewRepository[*todo.Todo, TodoTable](db)}
}

func (r *TodoRepository) Save(e *todo.Todo, ctx context.Context) error {
	var col TodoTable
	col = col.Transfer(e)

	if col.UpdatedAt != nil {
		log.Fatal("Updated at is not nil")
	}

	if err := r.db.WithContext(ctx).Save(&col).Error; err != nil {
		return err
	}
	e.SetPK(col.GetPK())

	return nil
}
