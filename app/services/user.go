package services

import (
	"errors"
	"go-chat/app/database"
	"go-chat/app/models"
	"go-chat/app/utils"
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

// GetUserAdmin - select
func (s UserAdmin) GetUserAdmin(email string, password string) (interface{}, error) {
	db := *database.Connection
	u := models.UserAdmin{}

	query := db.Where(&models.UserAdmin{MailAdress: email}).First(&u)
	count := query.RowsAffected

	log.Printf("vvvvvvvv %#v", u)
	log.Printf("ccccccccc %#v", count)

	if count == 0 {
		return nil, errors.New(utils.I18n.Translate("user_admin.error.not_found"))
	}

	if !utils.ComparePassword(u.HashedPassword, password) {
		return nil, errors.New(utils.I18n.Translate("user_admin.error.not_found"))
	}

	return u, nil
}

func isExistsMailAdress(s string) error {
	db := *database.Connection
	model := models.UserAdmin{MailAdress: s}

	query := db.Where(&model).First(&model)

	if query.RowsAffected != 0 {
		return errors.New(utils.I18n.Translate("validation_error.is_exists_mailadress"))
	}

	return nil
}
