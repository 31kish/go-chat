package db

import (
	"go-chat/app/models"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

// DB --
var DB *gorm.DB

// InitDB --
func InitDB() {
	db, err := gorm.Open("sqlite3", "./go-chat.db")

	if err != nil {
		revel.TRACE.Panicf("faild to connect database %v", err)
	}

	db.DB()
	DB = db

	autoMigrate()

	revel.TRACE.Println("connected to database")
}

func autoMigrate() {
	DB.AutoMigrate(models.UserAdmin{})
}
