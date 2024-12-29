package todo

import (
	"context"
	"sort"
	"time"

	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/infrastructure/common/db"
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

const TodoBucket = "TodoBucket"

func NewTodoRepository(i *do.Injector) (todo.TodoRepository, error) {
	bboltDB := do.MustInvoke[*bbolt.DB](i)
	if err := db.InitBuckets(bboltDB, TodoBucket); err != nil {
		return nil, err
	}
	return &todoRepository{
		db:         bboltDB,
		Repository: persistence.NewRepository[*todo.Todo, *TodoDBSchema](i, TodoBucket),
	}, nil
}

func (t *todoRepository) FindAllForNotification(ctx context.Context, dueDate time.Time) ([]*todo.Todo, error) {
	cols, err := db.GetAll[TodoDBSchema](t.db, TodoBucket, &db.GetAllOptions{OrderByDesc: true})
	if err != nil {
		return nil, err
	}

	cols = lo.Filter(cols, func(col TodoDBSchema, _ int) bool {
		return col.NotifiedAt == nil &&
			col.DueDate != nil && dueDate.UTC().After(lo.FromPtr(col.DueDate)) &&
			col.Status == enums.StatusTodo
	})

	sort.Slice(cols, func(i, j int) bool {
		return lo.FromPtr(cols[i].DueDate).Before(lo.FromPtr(cols[j].DueDate))
	})
	ret := lo.Map(cols, func(x TodoDBSchema, _ int) *todo.Todo { return x.ToEntity() })
	return ret, nil
}
