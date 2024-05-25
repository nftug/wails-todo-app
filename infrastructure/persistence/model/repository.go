package model

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/interfaces"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type Repository[TEntityPtr interfaces.Entity[TEntityPtr], TTablePtr EntityTable[TEntityPtr]] struct {
	db *gorm.DB
}

func NewRepository[TEntityPtr interfaces.Entity[TEntityPtr], TTablePtr EntityTable[TEntityPtr]](
	db *gorm.DB) *Repository[TEntityPtr, TTablePtr] {
	return &Repository[TEntityPtr, TTablePtr]{db}
}

func (r *Repository[TEntityPtr, TTablePtr]) Find(id uuid.UUID, ctx context.Context) (TEntityPtr, error) {
	col := *new(TTablePtr)
	if err := r.db.WithContext(ctx).Where("id = ?", id).Take(col).Error; err != nil {
		// レコードが見つからない場合は両方ともnilを返す
		return *new(TEntityPtr), filterNotFoundErr(err)
	}
	return col.ToEntity(), nil
}

func (r *Repository[TEntityPtr, TTablePtr]) Save(e TEntityPtr, ctx context.Context) error {
	col := *new(TTablePtr)
	col.Transfer(e)

	if err := r.db.WithContext(ctx).Save(col).Error; err != nil {
		return err
	}
	e.SetPK(col.GetPK())

	return nil
}

func (r *Repository[TEntityPtr, TTablePtr]) Delete(e TEntityPtr, ctx context.Context) error {
	if err := r.db.WithContext(ctx).Delete(*new(TTablePtr), e.PK()).Error; err != nil {
		return filterNotFoundErr(err)
	}
	return nil
}

func filterNotFoundErr(err error) error {
	return lo.Ternary(errors.Is(err, gorm.ErrRecordNotFound), nil, err)
}
