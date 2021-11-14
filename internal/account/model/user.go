package model

import "sports/pkg/db/data"

type User struct {
	ID       string `json:"id" gorm:"size:32;not null;primary_key;unique_index"`
	UserName string `json:"user_name" gorm:"not null;"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	data.BaseTimestampWithDelete
}

func (User) TableName() string {
	return "users"
}

func (u *User) IsExist() bool {
	return u.ID != ""
}
