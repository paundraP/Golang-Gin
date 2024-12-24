package migration

import (
	"log"
	"realtime-score/internal/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatalf("Error migrate the database: %v", err)
		return err
	}
	log.Println("Database migrations completed successfully!")
	return nil
}
