package common

import (
	"github.com/nftug/wails-todo-app/infrastructure/todo"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func NewDB(i *do.Injector) (*gorm.DB, error) {
	dialector := do.MustInvoke[gorm.Dialector](i)
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := autoMigrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

func autoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&todo.TodoDBSchema{}); err != nil {
		return err
	}
	return nil
}
