package testdb

import (
	"couponIssuanceSystem/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func NewTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(
		models.AllModels()...); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return db
}
