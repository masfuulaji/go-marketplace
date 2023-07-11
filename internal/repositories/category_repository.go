package repositories

import (
	"github.com/masfuulaji/go-marketplace/internal/config"
	"github.com/masfuulaji/go-marketplace/internal/models"
)

func CreateCategory(category models.Category) error {
    return config.DB.Create(&category).Error
}

func GetAllCategory() ([]models.Category, error) {
    var categories []models.Category
    return categories, config.DB.Find(&categories).Error
}

func GetCategoryById(categoryID string) (models.Category, error) {
    var category models.Category
    return category, config.DB.First(&category, categoryID).Error
}

func UpdateCategory(categoryID string, category models.Category) (models.Category, error) {
    var updatedCategory models.Category
    return updatedCategory, config.DB.Model(&updatedCategory).Where("id = ?", categoryID).Updates(category).Error
}

func DeleteCategory(categoryID string) error {
    return config.DB.Delete(&models.Category{}, categoryID).Error
}
