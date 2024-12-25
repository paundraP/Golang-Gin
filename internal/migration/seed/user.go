package seed

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"rest-api-go/internal/models"
	"rest-api-go/internal/pkg"

	"gorm.io/gorm"
)

func SeedingUser(db *gorm.DB) error {
	file, err := os.Open("internal/migration/data/users.json")
	if err != nil {
		log.Fatalf("Error opening seed data file: %v", err)
		return err
	}
	defer file.Close()

	var users []models.User
	if err := json.NewDecoder(file).Decode(&users); err != nil {
		log.Fatalf("Error decoding seed data: %v", err)
		return err
	}

	for i, user := range users {
		var existingUser models.User
		// cek kalo user udah ada di db
		if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
			log.Printf("Skipping user %d: email %s already exists", i, user.Email)
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Error checking user %d: %v", i, err)
			return err
		}

		hashedPassword, err := pkg.HashPassword(user.Password)
		if err != nil {
			log.Fatalf("Error hashing password for user %d: %v", i, err)
			return err
		}
		user.Password = hashedPassword

		if result := db.Create(&user); result.Error != nil {
			log.Printf("Error inserting user %d: %v", i, result.Error)
		} else {
			log.Printf("Inserted user %d: %s", i, user.Email)
		}
	}

	log.Println("Database seeding completed successfully!")
	return nil
}
