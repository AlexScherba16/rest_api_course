package postgres

import (
	"rest_api_course/config/database"
)

func NewPostgresConfig() database.DBConfig {
	return defaultPostgresConfig()
}

func defaultPostgresConfig() database.DBConfig {
	return database.DBConfig{
		Driver:              "postgres",
		Host:                "localhost",
		Port:                5432,
		DBName:              "postgres",
		SSLMode:             "disable",
		Username:            "postgres",
		Password:            "postgres",
		MaxOpenConnects:     10,
		MaxIdleConnects:     10,
		MaxLifeTimeConnects: 10,
	}
}

func envPostgresConfig() database.DBConfig {
	return database.DBConfig{}
}
