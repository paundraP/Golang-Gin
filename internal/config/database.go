package config

import (
	"fmt"
	"os"
	"realtime-score/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBNew() *gorm.DB {

	DBHost := os.Getenv("DBHOST")
	DBUser := os.Getenv("DBUSER")
	DBPassword := os.Getenv("DBPASSWORD")
	DBName := os.Getenv("DBNAME")
	DBPort := os.Getenv("DBPORT")

	DBDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		DBHost, DBUser, DBPassword, DBName, DBPort,
	)

	db, err := gorm.Open(postgres.Open(DBDSN), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("Failed to migrate database")
	}

	return db
}
