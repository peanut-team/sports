package model

import "sports/pkg/db/data"

type User struct {
	ID       int32  `json:"id" gorm:"size:32;not null;primary_key;unique_index"`
	UserType int32 `json:"user_type" gorm:"size:32;not null"`
	UserName string `json:"user_name" gorm:"not null;"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	data.BaseTimestamp
}

func (User) TableName() string {
	return "users"
}

func (u *User) IsExist() bool {
	return u.ID > 0
}
