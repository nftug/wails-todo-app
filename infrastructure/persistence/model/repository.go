package model

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/domain/shared/entity"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type Repository[TEntity entity.Entity[TEntity], TTable EntityTable[TEntity]] struct {
	db *gorm.DB
}

func NewRepository[TEntity entity.Entity[TEntity], TTable EntityTable[TEntity]](
	db *gorm.DB) *Repository[TEntity, TTable] {
	return &Repository[TEntity, TTable]{db}
}

func (r *Repository[TEntity, TTable]) Find(id uuid.UUID, ctx context.Context) (TEntity, error) {
	col := *new(TTable)
	if err := r.db.WithContext(ctx).Where("id = ?", id).Take(col).Error; err != nil {
		// レコードが見つからない場合は両方ともnilを返す
		return *new(TEntity), filterNotFoundErr(err)
	}
	return col.ToEntity(), nil
}

func (r *Repository[TEntity, TTable]) Save(e TEntity, ctx context.Context) error {
	col := *new(TTable)
	col.Transfer(e)

	if err := r.db.WithContext(ctx).Save(col).Error; err != nil {
		return err
	}
	e.SetPK(col.GetPK())

	return nil
}

func (r *Repository[TEntity, TTable]) Delete(e TEntity, ctx context.Context) error {
	if err := r.db.WithContext(ctx).Delete(*new(TTable), e.PK()).Error; err != nil {
		return filterNotFoundErr(err)
	}
	return nil
}

func filterNotFoundErr(err error) error {
	return lo.Ternary(errors.Is(err, gorm.ErrRecordNotFound), nil, err)
}
