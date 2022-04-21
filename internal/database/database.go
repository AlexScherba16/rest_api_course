package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"rest_api_course/config/database"
)

// NewDatabase - returns a pointer to a new database connection
func NewDatabase(cfg *database.DBConfig) (*gorm.DB, error) {

	connect := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.DBName,
		cfg.Password,
	)

	db, err := gorm.Open(cfg.DBName, connect)
	if err != nil {
		return db, err
	}

	if err := db.DB().Ping(); err != nil {
		return db, err
	}

	return db, nil

}
