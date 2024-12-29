package todo

import (
	"context"
	"strings"

	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/infrastructure/common/db"
	"github.com/samber/do"
	"github.com/samber/lo"
	"go.etcd.io/bbolt"
)

type todoQueryService struct {
	db *bbolt.DB
}

func NewTodoQueryService(i *do.Injector) (todo.TodoQueryService, error) {
	db := do.MustInvoke[*bbolt.DB](i)
	return &todoQueryService{db}, nil
}

func (qs *todoQueryService) Find(ctx context.Context, id int) (*todo.DetailsResponse, error) {
	col, err := db.Get[TodoDBSchema](qs.db, TodoBucket, id)
	if err != nil {
		return nil, err
	} else if col == nil {
		return nil, nil
	}
	return col.ToDetailsResponse(), nil
}

func (qs *todoQueryService) FindAll(ctx context.Context, query todo.Query) ([]*todo.ItemResponse, error) {
	items, err := db.GetAll[TodoDBSchema](qs.db, TodoBucket, &db.GetAllOptions{OrderByDesc: true})
	if err != nil {
		return nil, err
	}

	items = lo.Filter(items, func(item TodoDBSchema, _ int) bool {
		return matchesQuery(item, query)
	})
	ret := lo.Map(items, func(item TodoDBSchema, _ int) *todo.ItemResponse {
		return item.ToItemResponse()
	})

	return ret, nil
}

func matchesQuery(item TodoDBSchema, query todo.Query) bool {
	// Search条件
	if search := lo.FromPtr(query.Search); search != "" {
		if !strings.Contains(item.Title, search) &&
			!strings.Contains(lo.FromPtr(item.Description), search) {
			return false
		}
	}
	// Title条件
	if title := lo.FromPtr(query.Title); title != "" {
		if !strings.Contains(item.Title, title) {
			return false
		}
	}
	// Description条件
	if description := lo.FromPtr(query.Description); description != "" {
		if !strings.Contains(lo.FromPtr(item.Description), description) {
			return false
		}
	}
	// Status条件
	if status := lo.FromPtr(query.Status); status != "" {
		if item.Status != status {
			return false
		}
	}
	// Until条件
	if until := lo.FromPtr(query.Until); !until.IsZero() {
		if item.DueDate.After(until) {
			return false
		}
	}
	// After条件
	if after := lo.FromPtr(query.After); !after.IsZero() {
		if item.DueDate.Before(after) {
			return false
		}
	}
	// IsNotified条件
	if query.IsNotified != nil {
		if *query.IsNotified && item.NotifiedAt.IsZero() {
			return false
		}
		if !*query.IsNotified && !item.NotifiedAt.IsZero() {
			return false
		}
	}

	return true
}
