package todo

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/library/util"
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
	col := TodoDBSchema{}
	var item []byte

	if err := qs.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(TodoBucket))
		item = b.Get(util.Itob(id))
		return nil
	}); err != nil {
		return nil, err
	} else if item == nil {
		return nil, nil
	}

	if err := json.Unmarshal(item, &col); err != nil {
		return nil, fmt.Errorf("failed to unmarshal item: %w", err)
	}

	return col.ToDetailsResponse(), nil
}

func (qs *todoQueryService) FindAll(ctx context.Context, query todo.Query) ([]*todo.ItemResponse, error) {
	var items []TodoDBSchema

	err := qs.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(TodoBucket))

		c := bucket.Cursor()
		for k, v := c.Last(); k != nil; k, v = c.Prev() {
			var item TodoDBSchema
			if err := json.Unmarshal(v, &item); err != nil {
				return fmt.Errorf("failed to unmarshal item: %w", err)
			}
			items = append(items, item)
		}
		return nil
	})
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
