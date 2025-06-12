package routes

import (
	"github.com/gin-gonic/gin"
	"golang-clinic-app/controllers"
)

func RegisterAuthRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
}
