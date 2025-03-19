package main

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"server/internal/models"
)

func main() {
	// Connect to the database
	dsn := "host=localhost user=postgres password=secret dbname=server port=5442 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}


	// Run the migration to create the users table
	err = db.AutoMigrate(&models.Weather{})
	if err != nil {
		log.Fatalf("Error running migration: %v", err)
	}

	fmt.Println("Migration completed successfully.")
}
