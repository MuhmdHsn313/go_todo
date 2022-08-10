package datastore

import (
	"sample_rest_api/models"

	"gorm.io/gorm"
)

func Mirgate(db *gorm.DB) error {
	return db.AutoMigrate(&models.Todo{}, &models.User{}, &models.Session{})
}
