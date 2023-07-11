package models

import "time"

type Category struct {
    ID        int       `gorm:"primary_key"`
    Name      string    `gorm:"not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
    DeletedAt time.Time `gorm:"autoDeleteTime"`
}
