package ctr

import (
	"sports/pkg/db"

	"github.com/jinzhu/gorm"
)

var database *gorm.DB

// 数据库DB
func DB() *gorm.DB {
	return database.New()
}

func InitDatabase(config *db.Config) error {
	var err error
	database, err = db.NewDatabase(config)
	return err
}
