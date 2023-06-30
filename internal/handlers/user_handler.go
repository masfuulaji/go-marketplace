package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/masfuulaji/go-marketplace/internal/models"
	"github.com/masfuulaji/go-marketplace/internal/repositories"
)

func CreateUserHandler(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = repositories.CreateUser(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User created"})
}

func GetUserHandler(c *gin.Context) {
	userID := c.Param("userID")

	user, err := repositories.GetUserById(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func UpdateUserHandler(c *gin.Context) {
	var user models.User
	userID := c.Param("userID")

	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err = repositories.UpdateUser(userID, user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func DeleteUserHandler(c *gin.Context) {
	userID := c.Param("userID")

	err := repositories.DeleteUser(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted"})
}

func GetAllUserHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "All users"})
	// users, err := repositories.GetAllUser()
	// if err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }
	//
	// c.JSON(200, users)
}
