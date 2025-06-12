package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang-clinic-app/controllers"
	"golang-clinic-app/middleware"
)

func RegisterRoutes(router *gin.Engine) {
	// Public routes
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)

	// Simple test route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())

	// patient routes
	protected.GET("/protected", func(c *gin.Context) {
		userID := c.GetInt("userID")
		role := c.GetString("role")
		c.JSON(http.StatusOK, gin.H{
			"message": "Protected access granted",
			"userID":  userID,
			"role":    role,
		})
	})

	// Patient CRUD
	protected.POST("/patients", controllers.CreatePatient)
	protected.GET("/patients", controllers.GetPatients)
	protected.PUT("/patients/:id", controllers.UpdatePatient)
	protected.DELETE("/patients/:id", controllers.DeletePatient)

}
