package models

import (
	"time"
)

type Weather struct {
	ID                 uint      `gorm:"primaryKey"` // Auto-incrementing primary key
	CreatedAt          time.Time // Timestamp of record creation
	Pressure           float32   `gorm:"not null"`
	RelativeHumidity   float32   `gorm:"not null"`
	Temperature        float32   `gorm:"not null"`
	WindDirection      float32   `gorm:"not null"`
	WindSpeed          float32   `gorm:"not null"`
	CHP1               float32   `gorm:"not null"`
	DirectSun          float32   `gorm:"not null"`
	GlobalSun          float32   `gorm:"not null"`
	DiffuseSun         float32   `gorm:"not null"`
	RainFall           float32   `gorm:"not null"`
	AllDayIllumination float32   `gorm:"not null"`
	PM25               float32   `gorm:"not null"`
}
