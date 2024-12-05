package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`              // Auto-incrementing primary key
	Name      string         `gorm:"size:100;not null"`       // Name field with a maximum of 100 characters
	Email     string         `gorm:"size:150;unique;not null"`// Email field
	CreatedAt time.Time      // Timestamp of record creation
	UpdatedAt time.Time      // Timestamp of last update
	DeletedAt gorm.DeletedAt `gorm:"index"`                   // Soft delete column
}
