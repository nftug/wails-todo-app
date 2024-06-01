package model

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/interfaces"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type Repository[TEntityPtr interfaces.Entity[TEntityPtr], TSchema RepositorySchema[TEntityPtr, TSchema]] struct {
	db *gorm.DB
}

func NewRepository[TEntityPtr interfaces.Entity[TEntityPtr], TSchema RepositorySchema[TEntityPtr, TSchema]](
	db *gorm.DB) *Repository[TEntityPtr, TSchema] {
	return &Repository[TEntityPtr, TSchema]{db}
}

func (r *Repository[TEntityPtr, TSchema]) Find(ctx context.Context, id uuid.UUID) (TEntityPtr, error) {
	var col TSchema
	if err := r.db.WithContext(ctx).Where("id = ?", id).Take(&col).Error; err != nil {
		// レコードが見つからない場合は両方ともnilを返す
		return *new(TEntityPtr), filterNotFoundErr(err)
	}
	return col.ToEntity(), nil
}

func (r *Repository[TEntityPtr, TSchema]) Save(ctx context.Context, e TEntityPtr) error {
	var col TSchema
	col = col.Transfer(e)

	if err := r.db.WithContext(ctx).Save(&col).Error; err != nil {
		return err
	}
	e.SetPK(col.GetPK())

	return nil
}

func (r *Repository[TEntityPtr, TSchema]) Delete(ctx context.Context, e TEntityPtr) error {
	var col TSchema
	if err := r.db.WithContext(ctx).Delete(&col, e.PK()).Error; err != nil {
		return filterNotFoundErr(err)
	}
	return nil
}

func filterNotFoundErr(err error) error {
	return lo.Ternary(errors.Is(err, gorm.ErrRecordNotFound), nil, err)
}
