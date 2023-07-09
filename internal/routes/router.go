package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/masfuulaji/go-marketplace/internal/handlers"
	"github.com/masfuulaji/go-marketplace/internal/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	dashboard := router.Group("/")
	{
		dashboard.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	auth := router.Group("/auth")
	{
		auth.POST("/login", handlers.LoginHandler)
	}

	user := router.Group("/user")
    user.Use(middleware.AuthMiddleware())
	{
		user.GET("/", handlers.GetAllUserHandler)
		user.POST("/", handlers.CreateUserHandler)
		user.GET("/:userID", handlers.GetUserHandler)
		user.PUT("/:userID", handlers.UpdateUserHandler)
		user.DELETE("/:userID", handlers.DeleteUserHandler)
	}

	return router
}
