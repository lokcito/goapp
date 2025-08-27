package models

import (
    "time"

    "gorm.io/gorm"
)

var DB *gorm.DB

func SetDB(db *gorm.DB) {
    DB = db
}

// Robot model
type Robot struct {
    ID          uint           `gorm:"primaryKey" json:"id"`
    Nombre      string         `gorm:"size:255;not null" json:"nombre"`
    Descripcion string         `gorm:"type:text" json:"descripcion"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
