package models

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/jinzhu/gorm"
	"github.com/lisiur/daydayup/dbconn"
)

// User .
type User struct {
	gorm.Model
	Name           string `gorm:"unique;not null"`
	Email          string `gorm:"unique;not null"`
	PasswordDigest string `gorm:"not null"`
}

// All get all users.
func (*User) All() ([]User, error) {
	var users []User
	if err := dbconn.DB.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

// Create create user.
func (u *User) Create(input Register) error {
	var passwordDigestBytes = sha256.Sum256([]byte(input.Password))
	u.Name = input.Name
	u.Email = input.Email
	u.PasswordDigest = hex.EncodeToString(passwordDigestBytes[:])
	if err := dbconn.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

// Find .
func (u *User) Find(mp map[string]interface{}) error {
	if err := dbconn.DB.Where(mp).First(u).Error; err != nil {
		return err
	}
	return nil
}

// FindByID .
func (u *User) FindByID(id uint) error {
	if err := dbconn.DB.First(u, id).Error; err != nil {
		return err
	}
	return nil
}

// Authentication .
func (u *User) Authentication(input *Login) error {
	var passwordDigestBytes = sha256.Sum256([]byte(input.Password))
	var PasswordDigest = hex.EncodeToString(passwordDigestBytes[:])
	err := u.Find(map[string]interface{}{
		"name":            input.Name,
		"password_digest": PasswordDigest,
	})
	return err
}
