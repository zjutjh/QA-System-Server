package database

import (
	"QA-System-Server/app/models"
	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Question{},
		&models.NameMap{},
		&models.Submit{})
}
