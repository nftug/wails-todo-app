package db

import (
	"encoding/json"

	"github.com/nftug/wails-todo-app/library/util"
	"github.com/samber/lo"
	"go.etcd.io/bbolt"
)

func Get[TSchema any](db *bbolt.DB, bucketName string, id int) (*TSchema, error) {
	var col TSchema
	var v []byte

	if err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		v = b.Get(util.Itob(id))
		return nil
	}); err != nil {
		return nil, err
	} else if v == nil {
		return nil, nil
	}

	if err := json.Unmarshal(v, &col); err != nil {
		return nil, err
	}

	return &col, nil
}

type GetAllOptions struct {
	OrderByDesc bool
}

func GetAll[TSchema any](db *bbolt.DB, bucketName string, opt *GetAllOptions) ([]TSchema, error) {
	var items []TSchema

	appendParsed := func(data []byte) error {
		var item TSchema
		if err := json.Unmarshal(data, &item); err != nil {
			return err
		}
		items = append(items, item)
		return nil
	}

	err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))

		if lo.FromPtr(opt).OrderByDesc {
			c := b.Cursor()
			for k, v := c.Last(); k != nil; k, v = c.Prev() {
				if err := appendParsed(v); err != nil {
					return err
				}
			}
		} else {
			return b.ForEach(func(k, v []byte) error {
				if err := appendParsed(v); err != nil {
					return err
				}
				return nil
			})
		}

		return nil
	})

	return items, err
}
