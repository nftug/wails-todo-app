package todo

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/samber/do"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type todoQueryService struct {
	db *gorm.DB
}

func NewTodoQueryService(i *do.Injector) (todo.TodoQueryService, error) {
	db := do.MustInvoke[*gorm.DB](i)
	return &todoQueryService{db}, nil
}

func (qs *todoQueryService) Find(ctx context.Context, id uuid.UUID) (*todo.DetailResponse, error) {
	col := TodoDBSchema{}
	if err := qs.db.WithContext(ctx).Where("id = ?", id).Take(&col).Error; err != nil {
		// レコードが見つからない場合は両方ともnilを返す
		return nil, filterNotFoundErr(err)
	}
	return col.ToDetailResponse(), nil
}

func (qs *todoQueryService) FindAll(ctx context.Context, query todo.Query) ([]*todo.ItemResponse, error) {
	q := qs.db.WithContext(ctx)

	if lo.FromPtr(query.Search) != "" {
		q = q.Where("title LIKE %?%", query.Search).Or("description LIKE %?%", query.Search)
	}
	if lo.FromPtr(query.Title) != "" {
		q = q.Where("title LIKE %?%", query.Title)
	}
	if lo.FromPtr(query.Description) != "" {
		q = q.Where("description LIKE %?%", query.Description)
	}
	if lo.FromPtr(query.Status) != "" {
		q = q.Where("status = ?", query.Status)
	}
	if lo.IsNotEmpty(lo.FromPtr(query.Until)) {
		q = q.Where("due_date <= ?", query.Until.UTC())
	}
	if lo.IsNotEmpty(lo.FromPtr(query.After)) {
		q = q.Where("due_date >= ?", query.After.UTC())
	}
	if query.IsNotified != nil {
		if *query.IsNotified {
			q = q.Where("notified_at IS NOT NULL")
		} else {
			q = q.Where("notified_at IS NULL")
		}
	}

	var cols []TodoDBSchema
	if err := q.Order("created_at").Find(&cols).Error; err != nil {
		return nil, err
	}

	ret := lo.Map(cols, func(x TodoDBSchema, _ int) *todo.ItemResponse { return x.ToItemResponse() })
	return ret, nil
}

func filterNotFoundErr(err error) error {
	return lo.Ternary(errors.Is(err, gorm.ErrRecordNotFound), nil, err)
}
