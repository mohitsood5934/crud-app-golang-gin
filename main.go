package main

import (
	"log"
	"project/database"
	"project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize MongoDB connection
	client, err := database.InitMongoDB()
	if err != nil {
		log.Fatalf("Error initializing MongoDB: %v", err)
	}

	// Set the global client in the database package
	database.SetClient(client)

	// Create Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router)

	// Start the server
	log.Println("Server is running on http://localhost:4000")
	router.Run(":4000")
}
