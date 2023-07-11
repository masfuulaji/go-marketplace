package models

import "time"

type Product struct {
	ID          int       `gorm:"primary_key"`
	Name        string    `gorm:"not null"`
	Price       float64   `gorm:"not null"`
	Description string    `gorm:"size:255"`
	Stock       int       `gorm:"not null"`
	SellerID    int       `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	DeletedAt   time.Time `gorm:"autoDeleteTime"`
}
