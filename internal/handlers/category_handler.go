package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/masfuulaji/go-marketplace/internal/models"
	"github.com/masfuulaji/go-marketplace/internal/repositories"
)

func CreateCategoryHandler(c *gin.Context) {
    var category models.Category

    err := c.BindJSON(&category)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    err = repositories.CreateCategory(category)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"message": "Category created"})
}

func GetAllCategoryHandler(c *gin.Context) {
    categories, err := repositories.GetAllCategory()
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, categories)
}

func GetCategoryHandler(c *gin.Context) {
    categoryID := c.Param("categoryID")

    category, err := repositories.GetCategoryById(categoryID)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, category)
}

func UpdateCategoryHandler(c *gin.Context) {
    categoryID := c.Param("categoryID")

    var category models.Category

    err := c.BindJSON(&category)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    category, err = repositories.UpdateCategory(categoryID, category)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, category)
}

func DeleteCategoryHandler(c *gin.Context) {
    categoryID := c.Param("categoryID")

    err := repositories.DeleteCategory(categoryID)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"message": "Category deleted"})
}
