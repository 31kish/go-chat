package services

import (
	"go-chat/app/database"
	"go-chat/app/models"
	"go-chat/app/utils"
	"log"
)

// UserAdmin - buisiness logic
type UserAdmin struct {
	Base
}

// Create - Create UserAdmin
func (s UserAdmin) Create(u models.UserAdmin) (*models.UserAdmin, error) {
	db := *database.Connection

	if err := s.isExistsMailAdress(u.MailAdress); err != nil {
		return nil, err
	}

	result := db.Create(&u)

	if result.Error != nil {
		log.Printf("Error %#v", result.Error)
		return nil, result.Error
	}

	return &u, nil
}

// Update - Update UserAdmin
func (s UserAdmin) Update(id int, name string, mailadress string) error {
	db := *database.Connection

	u, err := s.GetByID(id)

	if err != nil {
		return err
	}

	err = s.isExistsMailAdressByID(id, mailadress)

	if err != nil {
		return err
	}

	update := models.UserAdmin{}

	if u.Name != name {
		update.Name = name
	}

	if u.MailAdress != mailadress {
		update.MailAdress = mailadress
	}

	result := db.Model(u).Update(update)

	if result.Error != nil {
		log.Printf("Error %#v", result.Error)
		return result.Error
	}

	return nil
}

// Delete - Delete UserAdmin
func (s UserAdmin) Delete(id int) error {
	db := *database.Connection

	user, err := s.GetByID(id)

	if err != nil {
		return err
	}

	result := db.Delete(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Get - Find UserAdmin with Email and Password
func (s UserAdmin) Get(email string, password string) (*models.UserAdmin, error) {
	db := *database.Connection
	u := models.UserAdmin{}

	query := db.Where(&models.UserAdmin{MailAdress: email}).First(&u)
	count := query.RowsAffected

	if count == 0 {
		return nil, s.notFound()
	}

	if !utils.ComparePassword(u.HashedPassword, password) {
		return nil, s.notFound()
	}

	return &u, nil
}

// GetByID - Find UserAdmin With ID
func (s UserAdmin) GetByID(id int) (*models.UserAdmin, error) {
	db := *database.Connection
	u := models.UserAdmin{}

	query := db.First(&u, id)
	count := query.RowsAffected

	if count == 0 {
		return nil, s.notFound()
	}

	return &u, nil
}

// GetAll - Return All UserAdmins
func (s UserAdmin) GetAll() ([]models.UserAdmin, error) {
	db := *database.Connection
	u := []models.UserAdmin{}

	query := db.Unscoped().Find(&u)
	count := query.RowsAffected

	if count == 0 {
		return nil, s.notFound()
	}

	return u, nil
}

func (s UserAdmin) isExistsMailAdress(email string) error {
	db := *database.Connection
	model := models.UserAdmin{MailAdress: email}

	query := db.Where(&model).First(&model)

	if query.RowsAffected != 0 {
		return s.existsMailAdress()
	}

	return nil
}

func (s UserAdmin) isExistsMailAdressByID(id int, email string) error {
	db := *database.Connection

	model := models.UserAdmin{}
	query := db.Not("id", id).Where(models.UserAdmin{MailAdress: email}).Find(&model)

	if query.RowsAffected != 0 {
		return s.existsMailAdress()
	}

	return nil
}
