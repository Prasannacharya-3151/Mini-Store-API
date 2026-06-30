package main

import (
	"log"
	"mini-store-api/config"
	"mini-store-api/models"
	"mini-store-api/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env files found, using system env vars")
	}

	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	log.Printf("server rnning on port %s", port)
	r.Run(":" + port)
}