package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/masfuulaji/go-marketplace/internal/models"
	"github.com/masfuulaji/go-marketplace/internal/repositories"
)

func CreateSellerHandler(c *gin.Context) {
	var seller models.Seller

	err := c.BindJSON(&seller)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = repositories.CreateSeller(seller)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Seller created"})
}

func GetAllSellerHandler(c *gin.Context) {
	sellers, err := repositories.GetAllSeller()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, sellers)
}

func GetSellerHandler(c *gin.Context) {
	sellerID := c.Param("sellerID")

	seller, err := repositories.GetSellerById(sellerID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, seller)
}

func UpdateSellerHandler(c *gin.Context) {
	sellerID := c.Param("sellerID")

	var seller models.Seller

	err := c.BindJSON(&seller)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	seller, err = repositories.UpdateSeller(sellerID, seller)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, seller)
}

func DeleteSellerHandler(c *gin.Context) {
    sellerID := c.Param("sellerID")

    err := repositories.DeleteSeller(sellerID)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"message": "Seller deleted"})
}
