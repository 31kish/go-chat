package models

import (
	"go-chat/app/utils"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

// User -
type User struct {
	gorm.Model
	MailAdress     string `gorm:"unique; size:100"`
	Name           string `gorm:"not null; size:18"`
	Password       string `gorm:"-"`
	HashedPassword []byte `gorm:"not null"`
}

// UserAdmin -
type UserAdmin struct {
	gorm.Model
	MailAdress     string `gorm:"unique; size:100"`
	Name           string `gorm:"not null; size:18"`
	Password       string `gorm:"-"`
	HashedPassword []byte `gorm:"not null"`
}

// Validate - UserAdmin
func (userAdmin *UserAdmin) Validate(v *revel.Validation) {
	v.Check(
		userAdmin.Name,
		revel.MaxSize{Max: 15},
		revel.MinSize{Min: 3},
	).Message(utils.I18n.Translate("validation_error.user_name"))

	v.Email(userAdmin.MailAdress).Message(utils.I18n.Translate("validation_error.email_format"))
	v.MaxSize(userAdmin.MailAdress, 50).Message(utils.I18n.Translate("validation_error.email_max_length"))
}
