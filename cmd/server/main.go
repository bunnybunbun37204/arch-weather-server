package main

import (
	"server/config"
	"server/internal/database"
	"server/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to the database
	db := database.ConnectDB(cfg)

	// Initialize Gin engine
	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r, db)

	// Start the server
	r.Run(":" + cfg.Port)
}
