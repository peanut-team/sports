package db

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Debug    bool   `json:"debug"`
}

func NewDatabase(cfg *Config) (*gorm.DB, error) {
	str := cfg.Username + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open("mysql", str)
	if err != nil {
		return nil, err
	}

	database.DB().SetMaxIdleConns(10)               // 连接池的空闲数大小
	database.DB().SetMaxOpenConns(100)              // 最大打开连接数
	database.DB().SetConnMaxLifetime(time.Hour * 6) // 连接最长存活时间
	database.SingularTable(true)                    // 禁用复数表名
	if cfg.Debug {
		database = database.Debug()
	}
	return database, nil
}
