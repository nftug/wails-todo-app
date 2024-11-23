package todo

import (
	"context"
	"time"

	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/infrastructure/persistence"
	"github.com/nftug/wails-todo-app/interfaces/enums"
	"github.com/samber/do"
	"github.com/samber/lo"
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

func (t *todoRepository) FindAllForNotification(ctx context.Context, dueDate time.Time) ([]*todo.Todo, error) {
	q := t.db.WithContext(ctx)

	q = q.Where("notified_at IS NULL")
	q = q.Where("due_date >= ?", dueDate.UTC())
	q = q.Where("status = ?", enums.StatusTodo)

	var cols []TodoDBSchema
	if err := q.Order("due_date").Find(&cols).Error; err != nil {
		return nil, err
	}

	ret := lo.Map(cols, func(x TodoDBSchema, _ int) *todo.Todo { return x.ToEntity() })
	return ret, nil
}
