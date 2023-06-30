package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/masfuulaji/go-marketplace/internal/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	dashboard := router.Group("/")
	{
		dashboard.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	user := router.Group("/user")
	{
		user.GET("/", handlers.GetAllUserHandler)
		user.POST("/", handlers.CreateUserHandler)
		user.GET("/:userID", handlers.GetUserHandler)
		user.PUT("/:userID", handlers.UpdateUserHandler)
		user.DELETE("/:userID", handlers.DeleteUserHandler)
	}

	return router
}
