package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/masfuulaji/go-marketplace/internal/models"
	"github.com/masfuulaji/go-marketplace/internal/repositories"
)

func CreateProductHandler(c *gin.Context) {
	var product models.Product

	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = repositories.CreateProduct(product)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Product created"})
}

func GetAllProductHandler(c *gin.Context) {
	products, err := repositories.GetAllProduct()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, products)
}

func GetProductHandler(c *gin.Context) {
	productID := c.Param("productID")

	product, err := repositories.GetProductById(productID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, product)
}

func UpdateProductHandler(c *gin.Context) {
	productID := c.Param("productID")

	var product models.Product

	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	product, err = repositories.UpdateProduct(productID, product)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, product)
}

func DeleteProductHandler(c *gin.Context) {
    productID := c.Param("productID")

    err := repositories.DeleteProduct(productID)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"message": "Product deleted"})
}
