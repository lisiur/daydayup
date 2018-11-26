package resolver

import (
	"context"
	"errors"
	"log"

	"github.com/lisiur/daydayup/util"

	"github.com/jinzhu/gorm"

	"github.com/lisiur/daydayup/models"
)

// Login .
func (r *MutationResolver) Login(ctx context.Context, input *models.Login) (models.LoginUser, error) {
	var user = models.User{}
	err := user.Authentication(input)

	if err != nil {
		log.Println(err.Error())
		if gorm.IsRecordNotFoundError(err) {
			err = errors.New("用户名或密码错误")
		}
		return models.LoginUser{}, err
	}

	tokenString, err := util.GenerateToken(user)
	if err != nil {
		log.Fatalln(err.Error())
	}
	var loginUser = models.LoginUser{
		UserID: user.ID,
		Token:  tokenString,
		User:   user,
	}

	if err = loginUser.CreateOrUpdate(); err != nil {
		return loginUser, err
	}

	return loginUser, nil
}
