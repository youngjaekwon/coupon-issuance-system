package db

import (
	"couponIssuanceSystem/internal/config"
	"couponIssuanceSystem/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

func Init() *gorm.DB {
	dsn := config.AppConfig.DatabaseURL
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	var db *gorm.DB
	var err error
	switch config.AppConfig.DBDriver {
	case "postgres":
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	default:
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	}
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	autoMigrate(db)

	return db
}

func autoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(models.AllModels()...); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
}
