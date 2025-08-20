package main

// This package serves as the entry point for the microserviceAuthGO application.
// It initializes the Echo web framework, loads environment variables, sets up middleware,
// connects to the database, initializes RabbitMQ, and starts the HTTP server.

import (
	"log"
	"microserviceAuthGO/src/db"
	"microserviceAuthGO/src/rabbitmq"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Load environment variables from the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set up Echo middleware for logging and panic recovery
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize the database connection
	db.Init()

	// Initialize RabbitMQ connection and channels
	rabbitmq.Init()

	// Start the HTTP server on port 8080
	if err := e.Start(":8080"); err != nil {
		// Close the database connection if the server fails to start
		db.Close()
		log.Fatalln(err)
	}

}
