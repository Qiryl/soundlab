package entity

import (
	"time"

	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

// User represent user entity.
type User struct {
	ID        string    `db:"user_id"`
	Email     string    `db:"email" validate:"required,email"`
	Name      string    `db:"username" validate:"required,gte=3"`
	FullName  string    `db:"full_name" validate:"required,gte=4"`
	Password  string    `db:"password" validate:"required,gte=8 lte=32"`
	Bio       *string   `db:"bio"`
	AvatarUrl *string   `db:"avatar_url"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// HashPassword hash user password using bcrypt.
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// ComparePasswords compare password.
func (u *User) ComparePasswords(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// PrepareCreate prepare user for register.
func (u *User) PrepareCreate() error {
	u.ID = xid.New().String()

	if err := u.HashPassword(); err != nil {
		return err
	}
	return nil
}

// SanitizePassword sanitizes user pwd.
func (u *User) SanitizePassword() {
	u.Password = ""
}
