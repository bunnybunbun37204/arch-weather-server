package repositories

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type WeatherRepository interface {
	Create(weather *models.Weather) error
	GetLatestWeather() (*models.Weather, error)
}

type weatherRepository struct {
	db *gorm.DB
}

// Create implements WeatherRepository.
func (w *weatherRepository) Create(weather *models.Weather) error {
	return w.db.Create(weather).Error
}

// GetLatestWeather implements WeatherRepository.
func (w *weatherRepository) GetLatestWeather() (*models.Weather, error) {
	var weather models.Weather

	err := w.db.Order("created_at DESC").First(&weather).Error

	return &weather, err
}

func NewWeatherRepository(db *gorm.DB) WeatherRepository {
	return &weatherRepository{db: db}
}
