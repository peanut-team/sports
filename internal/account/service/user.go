package service

import (
	"sports/internal/account/model"
	"sports/pkg/api/account"
	"sports/pkg/ctr"
)

func GetUser(username string) (*account.User, error) {
	user := &model.User{}
	ctr.DB().Where("user_name = ?", username).First(user)
	return &account.User{
		Username: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
