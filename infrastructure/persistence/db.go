package persistence

import (
	"log"

	"github.com/nftug/wails-todo-app/infrastructure/config"
	"github.com/nftug/wails-todo-app/infrastructure/todo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var instance *gorm.DB

func NewDB(lp *config.LocalPathService) *gorm.DB {
	if instance != nil {
		return instance
	}

	db, err := gorm.Open(sqlite.Open(lp.GetJoinedPath("todo.db")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&todo.TodoDBSchema{})

	instance = db
	return instance
}

func NewDBMock() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&todo.TodoDBSchema{})

	return db
}
