package models

import (
	"time"

	"github.com/revel/revel"
)

// User -
type User struct {
	ID             uint64 `gorm:"primary_key"`
	MailAdress     string `sql:"size:255"`
	Name           string `sql:"size:255"`
	Password       string `sql:"-"`
	HashedPassword []byte
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

// UserAdmin -
type UserAdmin struct {
	ID             uint64 `gorm:"primary_key"`
	MailAdress     string `sql:"size:255"`
	Name           string `sql:"size:255"`
	Password       string `sql:"-"`
	HashedPassword []byte
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

// Validate -
func (userAdmin *UserAdmin) Validate(v *revel.Validation) {
	v.Check(
		userAdmin.Name,
		revel.MaxSize{Max: 15},
		revel.MinSize{Min: 3},
	).Message("3文字以上、15文字以内で入力してください")

	v.Email(userAdmin.MailAdress).Message("メールアドレスの形式で入力してください")
}
