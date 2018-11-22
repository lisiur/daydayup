package models

import (
	"github.com/jinzhu/gorm"
)

// User .
type User struct {
	gorm.Model
	Name           string
	PasswordDigest string
	Email          string
}