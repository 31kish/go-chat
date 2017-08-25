package db

import (
	"go-chat/app/models"
	"log"

	"github.com/jinzhu/gorm"
	// sqlite -
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB -
var DB *gorm.DB

// InitDB -
func InitDB() {
	db, err := gorm.Open("sqlite3", "./go-chat.db")

	if err != nil {
		log.Panicf("faild to connect database %v", err)
	}

	db.DB()
	DB = db

	autoMigrate()

	log.Println("connected to database")
}

func autoMigrate() {
	DB.AutoMigrate(models.UserAdmin{}, models.User{})
}
