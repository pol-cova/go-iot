package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pol-cova/terrabyte/internal/api/routes"
	"log"
)

func main() {

	// Get env secret key
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Could not load .env file, default environment variables will be used")
	}

	server := gin.Default()

	server.LoadHTMLGlob("views/*")

	routes.RegisterRoutes(server)

	err = server.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
