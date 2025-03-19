package services

import (
	"encoding/json"
	"server/internal/models"
	"server/internal/repositories"
)

type WeatherService interface {
	CreateWeather(weather *models.Weather) error
	GetLatestWeather() (*models.Weather, error)
}

type weatherService struct {
	repo  repositories.WeatherRepository}

func NewWeatherService(repo repositories.WeatherRepository) WeatherService {
	return &weatherService{repo: repo}
}

func (s *weatherService) CreateWeather(weather *models.Weather) error {
	err := s.repo.Create(weather)
	if err != nil {
		return err
	}
	_, err = json.Marshal(weather)
	if err != nil {
		return err
	}

	return err
}

func (s *weatherService) GetLatestWeather() (*models.Weather, error) {

	weather, err := s.repo.GetLatestWeather()
	if err != nil {
		return nil, err
	}
	// Cache the result for future requests
	_, err = json.Marshal(weather)
	if err != nil {
		return nil, err
	}

	return weather, nil
}
