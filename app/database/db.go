package database

import (
	"fmt"
	"go-chat/app/models"
	"log"

	"github.com/revel/revel"

	"github.com/jinzhu/gorm"

	// sqlite -
	// _ "github.com/jinzhu/gorm/dialects/sqlite"

	// mysql -
	_ "github.com/go-sql-driver/mysql"
)

// Connection -
var Connection **gorm.DB

// Init -
func Init() {
	// db, err := gorm.Open("sqlite3", "./go-chat.db")
	connectionOption := getConnectionOption()
	db, err := gorm.Open("mysql", connectionOption)

	if err != nil {
		log.Panicf("faild to connect database %#v", err)
	}

	db.LogMode(revel.DevMode)

	db.AutoMigrate(models.UserAdmin{}, models.User{}, models.ChatEvent{})

	db.DB()
	Connection = &db

	log.Println("connected to database")
}

func getConnectionOption() string {
	user, found := revel.Config.String("db.user")
	if !found {
		panic("db user is required")
	}

	pass, found := revel.Config.String("db.password")
	if !found {
		panic("db password is required")
	}

	name, found := revel.Config.String("db.name")
	if !found {
		panic("db name is required")
	}

	host := revel.Config.StringDefault("db.host", "localhost")
	port := revel.Config.StringDefault("db.port", "3306")
	protocol := revel.Config.StringDefault("db.protocol", "tcp")

	revel.ERROR.Printf("%s:%s@%s([%s]:%s)/%s", user, pass, protocol, host, port, name)
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s", user, pass, protocol, host, port, name)
}
