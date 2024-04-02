package storage

import (
	"rinto11804/activity_point_backend/config"
	"rinto11804/activity_point_backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var defaultDB *gorm.DB

func ConnectDB(config *config.Config) {
	db, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})
	if err != nil {
		panic("DB Connection failed")
	}
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	db.Logger = logger.Default.LogMode(logger.Info)

	err = db.AutoMigrate(&models.Faculty{}, &models.Student{}, &models.Certificate{})
	if err != nil {
		panic("DB Migrations Failed")
	}

	defaultDB = db
}

func GetDB() *gorm.DB {
	return defaultDB
}
