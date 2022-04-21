package config

import (
	"rest_api_course/config/database"
	"rest_api_course/config/database/postgres"
)

type AppConfig struct {
	database.DBConfig
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		DBConfig: postgres.NewPostgresConfig(),
	}
}
