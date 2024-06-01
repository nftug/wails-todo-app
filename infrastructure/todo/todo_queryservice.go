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

func (qs *TodoQueryService) Find(ctx context.Context, id uuid.UUID) (*todo.DetailResponse, error) {
	col := TodoDBSchema{}
	if err := qs.db.WithContext(ctx).Where("id = ?", id).Take(&col).Error; err != nil {
		// レコードが見つからない場合は両方ともnilを返す
		return nil, filterNotFoundErr(err)
	}
	return col.ToDetailResponse(), nil
}

func (qs *TodoQueryService) FindAll(ctx context.Context, query todo.Query) ([]*todo.ItemResponse, error) {
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
