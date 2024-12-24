package cmd

import (
	"flag"
	"log"
	"realtime-score/internal/migration"

	"gorm.io/gorm"
)

func MigrateOrSeed(db *gorm.DB) {
	migrate := flag.Bool("migrate", false, "Run database migration")
	seed := flag.Bool("seed", false, "Run database seeder")

	flag.Parse()

	if *migrate {
		if err := migration.Migrate(db); err != nil {
			log.Fatalf("Error migrating database: %v", err)
		}
		log.Println("Database migration completed successfully.")
	}

	if *seed {
		if err := migration.Seeder(db); err != nil {
			log.Fatalf("Error seeding database: %v", err)
		}
		log.Println("Database seeding completed successfully.")
	}
}
