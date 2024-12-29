package db

import (
	"github.com/nftug/wails-todo-app/infrastructure/common/config"
	"github.com/samber/do"
	"go.etcd.io/bbolt"
)

func NewBBoltDB(i *do.Injector) (*bbolt.DB, error) {
	lp := do.MustInvoke[config.ConfigPathService](i)
	db, err := bbolt.Open(lp.GetJoinedPath("todo_bbolt.db"), 0600, nil)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitBuckets(db *bbolt.DB, bucketNames ...string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		for _, bucketName := range bucketNames {
			_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
			if err != nil {
				return err
			}
		}
		return nil
	})
}
