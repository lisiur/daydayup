package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lisiur/daydayup/dbconn"
)

// LoginUser .
type LoginUser struct {
	gorm.Model
	Token  string
	UserID uint
	User   User
}

// CreateOrUpdate .
func (u *LoginUser) CreateOrUpdate() error {
	var err error
	// user has login before.
	if err = dbconn.DB.Model(&u.User).Related(&u).Error; err == nil {
		if err = dbconn.DB.Model(&u).Update("token", u.Token).Error; err != nil {
			return err
		}
		return nil
	}
	// user has no login token before.
	if gorm.IsRecordNotFoundError(err) {
		// create user login token.
		err = dbconn.DB.Create(&u).Error
	}
	return err
}
