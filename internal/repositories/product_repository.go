package repositories

import (
	"github.com/masfuulaji/go-marketplace/internal/config"
	"github.com/masfuulaji/go-marketplace/internal/models"
)

func CreateProduct(seller models.Product) error {
    return config.DB.Create(&seller).Error
}

func GetAllProduct() ([]models.Product, error) {
    var products []models.Product
    return products, config.DB.Find(&products).Error
}

func GetProductById(productID string) (models.Product, error) {
    var product models.Product
    return product, config.DB.First(&product, productID).Error
}

func UpdateProduct(productID string, product models.Product) (models.Product, error) {
    var updatedProduct models.Product
    return updatedProduct, config.DB.Model(&updatedProduct).Where("id = ?", productID).Updates(product).Error
}

func DeleteProduct(productID string) error {
    return config.DB.Delete(&models.Product{}, productID).Error
}
