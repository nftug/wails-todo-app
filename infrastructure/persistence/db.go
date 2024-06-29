package persistence

import (
	"github.com/nftug/wails-todo-app/infrastructure/config"
	"github.com/nftug/wails-todo-app/infrastructure/todo"
	"github.com/samber/do"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var instance *gorm.DB

func NewDB(i *do.Injector) (*gorm.DB, error) {
	if instance != nil {
		return instance, nil
	}

	lp := do.MustInvoke[*config.LocalPathService](i)
	db, err := gorm.Open(sqlite.Open(lp.GetJoinedPath("todo.db")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&todo.TodoDBSchema{}); err != nil {
		return nil, err
	}

	instance = db
	return instance, nil
}

func NewDBMock(i *do.Injector) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&todo.TodoDBSchema{}); err != nil {
		return nil, err
	}

	return db, nil
}
