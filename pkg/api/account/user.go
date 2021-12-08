package account

import (
	"encoding/json"
	"sports/internal/account/model"
)

type UserType int

const (
	Coach_UserType UserType = iota
	Student_UserType
)

type User struct {
	ID       int32    `json:"id,omitempty"`                              // 用户 ID
	UserType UserType `json:"user_type"`                                 // 用户类型；0 为教练，1为学员
	Username string   `json:"username" example:"mick"`                   // 用户名
	Email    string   `json:"email" example:"123@ee.com"`                // 邮箱
	Password string   `json:"password,omitempty" example:"123@password"` // 密码
}

func NewUser(mUser *model.User) *User {
	return &User{
		ID:       mUser.ID,
		UserType: 0,
		Username: mUser.UserName,
		Email:    mUser.Email,
	}
}

func (u User) MarshalJSON() ([]byte, error) {
	tmpUser := u
	tmpUser.Password = ""
	// avoid circle invoke
	type Alias User
	return json.Marshal(Alias(tmpUser))
}

type UserItem struct {
	ID       int32    `json:"id"`        // 用户 ID
	UserType UserType `json:"user_type"` // 用户类型；0 为教练，1为学员
	Username string   `json:"username"`  // 用户名
	Email    string   `json:"email"`     // 邮箱
}

type LoginInfo struct {
	ID       int32    `form:"id"`        // 用户 ID
	Password string   `form:"password"`  // 密码
	UserType UserType `form:"user_type"` // 用户类型；0 为教练，1为学员
}

type AuthInfo struct {
	User  *UserItem `json:"user"`
	Token string    `json:"token"`
}
