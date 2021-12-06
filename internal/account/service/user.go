package service

import (
	"fmt"
	"sports/internal/account/model"
	"sports/pkg/api/account"
	"sports/pkg/ctr"
	"sports/pkg/errs"
	"sports/pkg/page"
	"sports/pkg/utils"
)

func GetUser(id int32) (*account.User, error) {
	user := &model.User{}
	err := ctr.DB().Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, errs.ErrorSystem.Params(err)
	}
	return &account.User{
		ID:       user.ID,
		Username: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func UserList(query page.Paginate) ([]*account.UserItem, page.Paginate, error) {
	userdb := ctr.DB().Model(&model.User{})
	var count int64
	err := userdb.Count(&count).Error //总行数
	if err != nil {
		return nil, query, errs.ErrorSystem.Params(err)
	}
	query.Total = count

	ul := []model.User{}
	err = userdb.Offset((query.Page - 1) * query.Page).Limit(query.PageSize).Find(&ul).Error //查询pageindex页的数据
	if err != nil {
		return nil, query, errs.ErrorSystem.Params(err)
	}
	users := make([]*account.UserItem, 0)
	for _, u := range ul {
		users = append(users, &account.UserItem{
			UserType: account.UserType(u.UserType),
			ID:       u.ID,
			Username: u.UserName,
			Email:    u.Email,
		})
	}
	return users, query, nil
}

func AddUser(user *account.User) (*account.User, error) {
	// hash pwd
	hpwd, err := utils.HashPWD(user.Password)
	if err != nil {
		return nil, fmt.Errorf("password encrypted failure")
	}
	mUser := &model.User{
		UserName: user.Username,
		Password: hpwd,
		Phone:    "",
		Email:    user.Email,
	}

	switch user.UserType {
	case account.Coach_UserType:
		mUser.UserType = 0
	case account.Student_UserType:
		mUser.UserType = 1
	default:
		mUser.UserType = 0
	}
	// create user
	err = ctr.DB().Create(mUser).Error
	if err != nil {
		return nil, errs.ErrorSystem.Params(err)
	}

	// clear pwd return
	user.Password = ""
	return user, nil
}

type UpdateUserOption string

const (
	Details_UpdateUserOption UpdateUserOption = "details"
	PWD_UpdateUserOption     UpdateUserOption = "pwd"
)

func UpdateUser(user *account.User, option UpdateUserOption) (*account.User, error) {
	// get user
	mUser := &model.User{}
	err := ctr.DB().Where("id = ?", user.ID).First(mUser).Error
	if err != nil {
		return nil, errs.ErrorParams.Params(err)
	}

	switch option {
	case Details_UpdateUserOption:
		if err := syncUserDetails(user, mUser); err != nil {
			return nil, errs.ErrorParams.Params(err)
		}
	case PWD_UpdateUserOption:
		if err := syncUserPWD(user, mUser); err != nil {
			return nil, errs.ErrorParams.Params(err)
		}
	default:
		return nil, errs.ErrorParams.Params(fmt.Errorf("invalid update option: %s", option))
	}

	// update user
	err = ctr.DB().Save(mUser).Error
	if err != nil {
		return nil, errs.ErrorSystem.Params(fmt.Errorf("update user failed: %v", err))
	}

	return account.NewUser(mUser), nil
}

func syncUserDetails(user *account.User, mUser *model.User) error {
	// todo check email format
	// sync email
	mUser.Email = user.Email

	return nil
}

func syncUserPWD(user *account.User, mUser *model.User) error {
	// hash pwd
	hpwd, err := utils.HashPWD(user.Password)
	if err != nil {
		return fmt.Errorf("password encrypted failure")
	}

	// sync PWD
	mUser.Password = hpwd

	return nil
}

func DeleteUser(id int32) error {

	mUser := &model.User{
		ID: id,
	}

	err := ctr.DB().Delete(mUser).Error
	if err != nil {
		return errs.ErrorParams.Params(fmt.Errorf("delete user id (%s) faild: %v", mUser.ID, err))
	}
	return nil
}
