package routes

import (
	"mini-store-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/signup", handlers.Signup)
		api.POST("/login", handlers.Login)
	}
}