package model

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/interfaces"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type Repository[TEntityPtr interfaces.Entity[TEntityPtr], TTable EntityTable[TEntityPtr, TTable]] struct {
	db *gorm.DB
}

func NewRepository[TEntityPtr interfaces.Entity[TEntityPtr], TTable EntityTable[TEntityPtr, TTable]](
	db *gorm.DB) *Repository[TEntityPtr, TTable] {
	return &Repository[TEntityPtr, TTable]{db}
}

func (r *Repository[TEntityPtr, TTable]) Find(ctx context.Context, id uuid.UUID) (TEntityPtr, error) {
	var col TTable
	if err := r.db.WithContext(ctx).Where("id = ?", id).Take(&col).Error; err != nil {
		// レコードが見つからない場合は両方ともnilを返す
		return *new(TEntityPtr), filterNotFoundErr(err)
	}
	return col.ToEntity(), nil
}

func (r *Repository[TEntityPtr, TTable]) Save(ctx context.Context, e TEntityPtr) error {
	var col TTable
	col = col.Transfer(e)

	if err := r.db.WithContext(ctx).Save(&col).Error; err != nil {
		return err
	}
	e.SetPK(col.GetPK())

	return nil
}

func (r *Repository[TEntityPtr, TTable]) Delete(ctx context.Context, e TEntityPtr) error {
	var col TTable
	if err := r.db.WithContext(ctx).Delete(&col, e.PK()).Error; err != nil {
		return filterNotFoundErr(err)
	}
	return nil
}

func filterNotFoundErr(err error) error {
	return lo.Ternary(errors.Is(err, gorm.ErrRecordNotFound), nil, err)
}
