package todo

import (
	"context"
	"encoding/json"
	"sort"
	"time"

	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/infrastructure/persistence"
	"github.com/nftug/wails-todo-app/shared/enums"
	"github.com/samber/do"
	"github.com/samber/lo"
	"go.etcd.io/bbolt"
)

type todoRepository struct {
	db *bbolt.DB
	*persistence.Repository[*todo.Todo, *TodoDBSchema]
}

func NewTodoRepository(i *do.Injector) (todo.TodoRepository, error) {
	return &todoRepository{
		db:         do.MustInvoke[*bbolt.DB](i),
		Repository: persistence.NewRepository[*todo.Todo, *TodoDBSchema](i, TodoBucket),
	}, nil
}

func (t *todoRepository) FindAllForNotification(ctx context.Context, dueDate time.Time) ([]*todo.Todo, error) {
	var cols []TodoDBSchema

	if err := t.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(TodoBucket))
		return b.ForEach(func(k, v []byte) error {
			var col TodoDBSchema
			if err := json.Unmarshal(v, &col); err != nil {
				return err
			}

			if col.NotifiedAt == nil &&
				col.DueDate != nil && dueDate.UTC().After(lo.FromPtr(col.DueDate)) &&
				col.Status == enums.StatusTodo {
				cols = append(cols, col)
			}

			return nil
		})
	}); err != nil {
		return nil, err
	}

	sort.Slice(cols, func(i, j int) bool {
		return lo.FromPtr(cols[i].DueDate).Before(lo.FromPtr(cols[j].DueDate))
	})
	ret := lo.Map(cols, func(x TodoDBSchema, _ int) *todo.Todo { return x.ToEntity() })
	return ret, nil
}
