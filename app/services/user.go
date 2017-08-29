package services

import (
	"errors"
	"go-chat/app/database"
	"go-chat/app/models"
	"log"
)

// UserAdmin - buisiness logic
type UserAdmin struct {
}

// Save - insert
func (s UserAdmin) Save(i models.UserAdmin) (interface{}, error) {
	db := *database.Connection

	if err := isExistsMailAdress(i.MailAdress); err != nil {
		return nil, err
	}

	result := db.Create(&i)

	if result.Error != nil {
		log.Printf("Error %#v", result.Error)
	}

	return result.Value, result.Error
}

func isExistsMailAdress(s string) error {
	db := *database.Connection
	result := db.Where(models.UserAdmin{MailAdress: s}).First(&models.UserAdmin{})

	if result.Value != nil {
		return errors.New("すでに登録済みのメールアドレスです。")
	}

	return nil
}
