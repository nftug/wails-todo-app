package todo

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type TodoQueryService struct {
	db *gorm.DB
}

func NewTodoQueryService(db *gorm.DB) todo.TodoQueryService {
	return &TodoQueryService{db}
}

func (qs *TodoQueryService) Find(id uuid.UUID, ctx context.Context) (*todo.DetailResponse, error) {
	col := TodoTable{}
	if err := qs.db.WithContext(ctx).Where("id = ?", id).Take(col).Error; err != nil {
		// レコードが見つからない場合は両方ともnilを返す
		return nil, filterNotFoundErr(err)
	}
	return col.ToDetailResponse(), nil
}

func (qs *TodoQueryService) FindAll(query todo.Query, ctx context.Context) ([]*todo.ItemResponse, error) {
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
		q = q.Where("due_date <= ?", query.Until)
	}

	var cols []TodoTable
	if err := q.Order("created_at").Find(&cols).Error; err != nil {
		return nil, err
	}

	ret := lo.Map(cols, func(x TodoTable, _ int) *todo.ItemResponse { return x.ToItemResponse() })
	return ret, nil
}

func filterNotFoundErr(err error) error {
	return lo.Ternary(errors.Is(err, gorm.ErrRecordNotFound), nil, err)
}
