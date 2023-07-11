package models

import "time"

type Seller struct {
    ID        int       `gorm:"primary_key"`
    UserID    int       `gorm:"not null"`
    Name      string    `gorm:"not null"`
    Email     string    `gorm:"not null"`
    Address   string    `gorm:"size:255"`
    Phone     string    `gorm:"size:100"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
    DeletedAt time.Time `gorm:"autoDeleteTime"`
}
