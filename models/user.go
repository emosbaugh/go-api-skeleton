package models

import (
	"time"

	passlib "gopkg.in/hlandau/passlib.v1"
)

type User struct {
	Email             string
	PasswordHash      string    `db:"password"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
	PasswordUpdatedAt time.Time `db:"password_updated_at"`
}

func (u *User) VerifyPassword(password string) error {
	return passlib.VerifyNoUpgrade(password, u.PasswordHash)
}
