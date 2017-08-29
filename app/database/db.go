package database

import (
	"go-chat/app/models"
	"log"

	"github.com/revel/revel"

	"github.com/jinzhu/gorm"
	// sqlite -
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Connection -
var Connection **gorm.DB

// Init -
func Init() {
	db, err := gorm.Open("sqlite3", "./go-chat.db")

	if err != nil {
		log.Panicf("faild to connect database %#v", err)
	}

	db.LogMode(revel.DevMode)
	db.DB()
	Connection = &db

	autoMigrate()

	log.Println("connected to database")
}

func autoMigrate() {
	conn := *Connection
	conn.AutoMigrate(models.UserAdmin{}, models.User{})
}

type DataServiceBase struct {
	*gorm.DB
}
