package database

import (
	"log"
	"sync"

	"github.com/Bash360/go-template/internal/modules/users/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func Connect() *gorm.DB {
	once.Do(func() {
		database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

		if err != nil {
			log.Fatal("Database error ", err.Error())
		}
		log.Println("database connected successfully")

		database.AutoMigrate(&model.User{})
		DB = database

	})

	return DB
}
