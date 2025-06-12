package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"golang-clinic-app/config"
	"golang-clinic-app/routes"
)

func main() {
	// Create Gin router
	r := gin.Default()

	// Connect to DB
	config.ConnectToDB()

	// Register routes
	routes.RegisterRoutes(r)

	// Get port from environment (Render sets this)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local dev
	}

	// Start the server
	fmt.Println("Server running at http://localhost:" + port)
	r.Run(":" + port)
}
