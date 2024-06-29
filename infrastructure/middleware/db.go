package middleware

import (
	"github.com/nftug/wails-todo-app/infrastructure/todo"
	"github.com/samber/do"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB(i *do.Injector) (*gorm.DB, error) {
	lp := do.MustInvoke[LocalPathService](i)
	db, err := gorm.Open(sqlite.Open(lp.GetJoinedPath("todo.db")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := autoMigrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

func NewDBMock(i *do.Injector) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
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
