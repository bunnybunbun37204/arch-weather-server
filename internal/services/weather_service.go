package services

import (
	"server/internal/models"
	"server/internal/repositories"
)

type WeatherService interface {
	CreateWeather(weather *models.Weather) error
	GetLatestWeather() (*models.Weather, error)
}

type weatherService struct {
	repo repositories.WeatherRepository
}

func NewWeatherService(repo repositories.WeatherRepository) WeatherService {
	return &weatherService{repo: repo}
}

func (s *weatherService) CreateWeather(weather *models.Weather) error {
	return s.repo.Create(weather)
}

func (s *weatherService) GetLatestWeather() (*models.Weather, error) {
	return s.repo.GetLatestWeather()
}
