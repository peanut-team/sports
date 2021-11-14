package base

import "time"

const (
	ConfigPath    = "./config/config.yaml"
	EnvProduction = "production"
	ServerPort    = 8080
)

type Time time.Time
