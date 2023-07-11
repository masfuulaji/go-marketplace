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

	login := router.Group("/auth")
	{
		login.POST("/login", handlers.LoginHandler)
	}

    logout := router.Group("/auth")
    logout.Use(middleware.AuthMiddleware())
    {
        logout.POST("/logout", handlers.LogoutHandler)
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

    category := router.Group("/category")
    {
        category.GET("/", handlers.GetAllCategoryHandler)
        category.POST("/", handlers.CreateCategoryHandler)
        category.GET("/:categoryID", handlers.GetCategoryHandler)
        category.PUT("/:categoryID", handlers.UpdateCategoryHandler)
        category.DELETE("/:categoryID", handlers.DeleteCategoryHandler)
    }

    product := router.Group("/product")
    {
        product.GET("/", handlers.GetAllProductHandler)
        product.POST("/", handlers.CreateProductHandler)
        product.GET("/:productID", handlers.GetProductHandler)
        product.PUT("/:productID", handlers.UpdateProductHandler)
        product.DELETE("/:productID", handlers.DeleteProductHandler)
    }

    seller := router.Group("/seller")
    {
        seller.GET("/", handlers.GetAllSellerHandler)
        seller.POST("/", handlers.CreateSellerHandler)
        seller.GET("/:sellerID", handlers.GetSellerHandler)
        seller.PUT("/:sellerID", handlers.UpdateSellerHandler)
        seller.DELETE("/:sellerID", handlers.DeleteSellerHandler)
    }

	return router
}
