package models

import (
	"time"
)

type Weather struct {
	ID                 uint      `gorm:"primaryKey"` // Auto-incrementing primary key
	CreatedAt          time.Time // Timestamp of record creation
	Pressure           float32   `json:"pressure" gorm:"not null"`
	RelativeHumidity   float32   `json:"relative_humidity" gorm:"not null"`
	Temperature        float32   `json:"temperature" gorm:"not null"`
	WindDirection      float32   `json:"wind_direction" gorm:"not null"`
	WindSpeed          float32   `json:"wind_speed" gorm:"not null"`
	CHP1               float32   `json:"chp1" gorm:"not null"`
	DirectSun          float32   `json:"direct_sun" gorm:"not null"`
	GlobalSun          float32   `json:"global_sun" gorm:"not null"`
	DiffuseSun         float32   `json:"diffuse_sun" gorm:"not null"`
	RainFall           float32   `json:"rain_fall" gorm:"not null"`
	AllDayIllumination float32   `json:"all_day_illumination" gorm:"not null"`
	PM25               float32   `json:"pm25" gorm:"not null"`
}
