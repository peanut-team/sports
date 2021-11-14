package config

import (
	"sports/pkg/common/server"
	"sports/pkg/db"
)

type Config struct {
	Server server.Config `json:"server"`
	Mysql  db.Config     `json:"mysql"`

	Debug bool `json:"debug"`
	// Secret for jwt token
	Secret string `json:"secret"`
}

var cfg *Config

func GetConfig() Config {
	return *cfg
}

func InitConfig(c Config) {
	cfg = &c
}
