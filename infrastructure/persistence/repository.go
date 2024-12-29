package persistence

import (
	"context"
	"encoding/json"

	"github.com/nftug/wails-todo-app/infrastructure/common/db"
	"github.com/nftug/wails-todo-app/library/util"
	"github.com/nftug/wails-todo-app/shared/interfaces"
	"github.com/samber/do"
	"github.com/samber/lo"
	"go.etcd.io/bbolt"
)

type Repository[
	TEntityPtr interfaces.Entity[TEntityPtr],
	TSchema RepositorySchema[TEntityPtr, TSchema],
] struct {
	db         *bbolt.DB
	bucketName string
}

func NewRepository[
	TEntityPtr interfaces.Entity[TEntityPtr],
	TSchema RepositorySchema[TEntityPtr, TSchema],
](
	i *do.Injector, bucketName string) *Repository[TEntityPtr, TSchema] {
	return &Repository[TEntityPtr, TSchema]{
		db:         do.MustInvoke[*bbolt.DB](i),
		bucketName: bucketName,
	}
}

func (r *Repository[TEntityPtr, TSchema]) Find(ctx context.Context, id int) (TEntityPtr, error) {
	col, err := db.Get[TSchema](r.db, r.bucketName, id)
	if err != nil {
		return *new(TEntityPtr), err
	} else if col == nil {
		return *new(TEntityPtr), nil
	}
	return lo.FromPtr(col).ToEntity(), nil
}

func (r *Repository[TEntityPtr, TSchema]) Save(ctx context.Context, e TEntityPtr) error {
	var col TSchema

	if err := r.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(r.bucketName))
		if e.ID() == 0 {
			newID, _ := b.NextSequence()
			e.SetID(int(newID))
		}

		item, err := json.Marshal(col.Transfer(e))
		if err != nil {
			return err
		}

		id := e.ID()
		if err := b.Put(util.Itob(id), item); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *Repository[TEntityPtr, TSchema]) Delete(ctx context.Context, id int) error {
	return r.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(r.bucketName))
		if err := b.Delete(util.Itob(id)); err != nil {
			return err
		}
		return nil
	})
}
