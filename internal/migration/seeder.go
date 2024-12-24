package migration

import (
	"realtime-score/internal/migration/seed"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := seed.SeedingUser(db); err != nil {
		return err
	}
	return nil
}
