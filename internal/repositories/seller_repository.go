package repositories

import (
	"github.com/masfuulaji/go-marketplace/internal/config"
	"github.com/masfuulaji/go-marketplace/internal/models"
)

func CreateSeller(seller models.Seller) error {
    return config.DB.Create(&seller).Error
}

func GetAllSeller() ([]models.Seller, error) {
    var sellers []models.Seller
    return sellers, config.DB.Find(&sellers).Error
}

func GetSellerById(sellerID string) (models.Seller, error) {
    var seller models.Seller
    return seller, config.DB.First(&seller, sellerID).Error
}

func UpdateSeller(sellerID string, seller models.Seller) (models.Seller, error) {
    var updatedSeller models.Seller
    return updatedSeller, config.DB.Model(&updatedSeller).Where("id = ?", sellerID).Updates(seller).Error
}

func DeleteSeller(sellerID string) error {
    return config.DB.Delete(&models.Seller{}, sellerID).Error
}
