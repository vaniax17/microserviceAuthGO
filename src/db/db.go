package db

// This package handles the database connection and migration for the application.
// It provides functions to initialize and close the PostgreSQL database connection
// using GORM.

import (
	"fmt"
	"log"
	"microserviceAuthGO/src/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB // Global variable to store the GORM database connection.
	err error    // Global variable to store any error that occurs during database operations.
)

// Init initializes the database connection and performs auto-migration.
func Init() {
	// Retrieve database connection details from environment variables.
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	// Construct the Data Source Name (DSN) for connecting to the PostgreSQL database.
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)
	// Open a connection to the PostgreSQL database using GORM.
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to database")
	// Automatically migrate the database schema based on defined GORM models.
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migrated")

}

// Close closes the database connection.
func Close() {
	sqlDB, _ := DB.DB()
	// Close the database connection.
	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("Failed to close database: %v", err)
	}
}
