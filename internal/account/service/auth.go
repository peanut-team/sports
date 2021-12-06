package service

import (
	"fmt"
	"sports/internal/account/model"
	"sports/pkg/api/account"
	"sports/pkg/ctr"
	"sports/pkg/errs"
	"sports/pkg/utils"
)

func Login(user *account.LoginInfo) (*account.AuthInfo, error) {
	// get user
	mu := &model.User{}
	err := ctr.DB().Where("id = ?", user.ID).First(mu).Error
	if err != nil {
		return nil, err
	}

	// check pwd
	validPwd, _ := utils.ComparePasswords(mu.Password, user.Password)
	if !validPwd {
		return nil, errs.ErrorParams.Params(fmt.Errorf("id or password is invalid"))
	}

	return &account.AuthInfo{
		User: &account.UserItem{
			ID:       mu.ID,
			UserType: 0,
			Username: mu.UserName,
			Email:    mu.Email,
		},
	}, nil
}
