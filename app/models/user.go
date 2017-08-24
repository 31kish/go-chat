package models

import (
	"time"
)

// User --
type User struct {
	ID         uint64 `gorm:"primary_key"`
	MailAdress string `sql:"size:255" validate:"email"`
	Name       string `sql:"size255" validate:"min=1"`
	Password   string `sql:"size255"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

// UserAdmin --
type UserAdmin struct {
	ID         uint64 `gorm:"primary_key"`
	MailAdress string `sql:"size:255" validate:"email"`
	Name       string `sql:"size255" validate:"min=1"`
	Password   string `sql:"size255"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
