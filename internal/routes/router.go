package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/masfuulaji/go-marketplace/internal/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("/create", handlers.CreateUserHandler)
		user.GET("/:userID", handlers.GetUserHandler)
		user.PUT("/:userID", handlers.UpdateUserHandler)
		user.DELETE("/:userID", handlers.DeleteUserHandler)
		user.GET("/", handlers.GetAllUserHandler)
	}

	return router
}
