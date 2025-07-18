package main

import (
	"server/config"
	"server/internal/database"
	"server/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to the database
	db := database.ConnectDB(cfg)

	// Initialize Gin engine
	r := gin.Default()

	// Enable CORS for all origins
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Register routes
	routes.RegisterRoutes(r, db)

	// Start the server
	r.Run(":" + cfg.Port)
}
