package main

import (
	"fmt"

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

	// Start the server
	port := ":8080"
	fmt.Println("Server running at http://localhost" + port)
	r.Run(port)
}
