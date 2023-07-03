package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/masfuulaji/go-marketplace/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

func loginHandler(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	result, err := repositories.GetUserByEmail(email)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "login success"})
}
