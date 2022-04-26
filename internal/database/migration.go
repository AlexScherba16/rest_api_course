package database

import (
	"github.com/jinzhu/gorm"
	"rest_api_course/internal/comment"
)

//MigrateDB -migrates our database and creates out comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result != nil {
		return result.Error
	}
	return nil
}
