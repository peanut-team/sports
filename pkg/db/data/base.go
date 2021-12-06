package data

import (
	"sports/pkg/common/base"
)

type AutoID struct {
	ID uint `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
}

type SingleTimestamp struct {
	CreateTime       *base.Time `json:"create_time" gorm:"type:datetime;not null;"`
	LastModifiedTime *base.Time `json:"last_modified_time" gorm:"type:datetime;"`
}

type BaseTimestamp struct {
	CreateTime       *base.Time `json:"create_time" gorm:"type:datetime;not null;"`
	CreateBy         string    `json:"create_by" gorm:"size:20;not null;"`
	LastModifiedTime *base.Time `json:"last_modified_time" gorm:"type:datetime;"`
	LastModifiedBy   string    `json:"last_modified_by" gorm:"size:20;"`
}

type BaseTimestampWithDelete struct {
	BaseTimestamp
	//	DeletedAt *time.Time `gorm:"column:deleted_at;index"`
	Deleted bool `json:"deleted" gorm:"not null;default:0"`
}
