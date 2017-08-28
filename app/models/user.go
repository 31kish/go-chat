package models

import (
	"time"

	"github.com/revel/revel"
)

// modelBase - Base definition
type modelBase struct {
	ID        uint `gorm:"primary_key; AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// User -
type User struct {
	modelBase
	MailAdress     string `gorm:"unique; size:100"`
	Name           string `gorm:"not null; size:18"`
	Password       string `gorm:"-"`
	HashedPassword []byte `gorm:"not null"`
}

// UserAdmin -
type UserAdmin struct {
	modelBase
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
	).Message("3文字以上、15文字以内で入力してください")

	v.Email(userAdmin.MailAdress).Message("メールアドレスの形式で入力してください")
	v.MaxSize(userAdmin.MailAdress, 50).Message("50文字以内で入力してください")
}
