package services

import (
	"go-chat/app/database"
	"go-chat/app/models"
	"go-chat/app/utils"
	"log"
)

// User -
type User struct {
	Base
}

// Create -
func (s User) Create(u models.User) (*models.User, error) {
	db := *database.Connection

	if err := s.isExistsMailAdress(u.MailAdress); err != nil {
		return nil, err
	}

	result := db.Create(&u)

	if result.Error != nil {
		log.Printf("%#v", result.Error)
		return nil, result.Error
	}

	return &u, nil
}

// Get -
func (s User) Get(email string, password string) (*models.User, error) {
	db := *database.Connection
	u := models.User{}

	query := db.Where(&models.User{MailAdress: email}).First(&u)
	count := query.RowsAffected

	if count == 0 {
		return nil, s.notFound()
	}

	if !utils.ComparePassword(u.HashedPassword, password) {
		return nil, s.notFound()
	}

	return &u, nil
}

// GetByID -
func (s User) GetByID(id int) (*models.User, error) {
	db := *database.Connection
	u := models.User{}

	query := db.First(&u, id)
	count := query.RowsAffected

	if count == 0 {
		return nil, s.notFound()
	}

	return &u, nil
}

// GetAll -
func (s User) GetAll() ([]models.User, error) {
	db := *database.Connection
	u := []models.User{}

	query := db.Unscoped().Find(&u)
	count := query.RowsAffected

	if count == 0 {
		return nil, s.notFound()
	}

	return u, nil
}

func (s User) Update(id int, name string, mailadress string) error {
	db := *database.Connection

	u, err := s.GetByID(id)

	if err != nil {
		return err
	}

	err = s.isExistsMailAdressByID(id, mailadress)

	if err != nil {
		return err
	}

	update := models.User{}

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

// Delete -
func (s User) Delete(id int) error {
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

func (s User) isExistsMailAdress(email string) error {
	db := *database.Connection
	model := models.User{MailAdress: email}

	query := db.Where(&model).First(&model)

	if query.RowsAffected != 0 {
		return s.existsMailAdress()
	}

	return nil
}

func (s User) isExistsMailAdressByID(id int, email string) error {
	db := *database.Connection

	model := models.User{}
	query := db.Not("id", id).Where(models.User{MailAdress: email}).Find(&model)

	if query.RowsAffected != 0 {
		return s.existsMailAdress()
	}

	return nil
}
