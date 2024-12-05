package services

import (
	"encoding/json"
	"server/internal/cache"
	"server/internal/models"
	"server/internal/repositories"
	"time"
)

type WeatherService interface {
	CreateWeather(weather *models.Weather) error
	GetLatestWeather() (*models.Weather, error)
}

type weatherService struct {
	repo  repositories.WeatherRepository
	cache *cache.RedisCache
}

func NewWeatherService(repo repositories.WeatherRepository, cache *cache.RedisCache) WeatherService {
	return &weatherService{repo: repo, cache: cache}
}

func (s *weatherService) CreateWeather(weather *models.Weather) error {
	err := s.repo.Create(weather)
	if err != nil {
		return err
	}
	cacheKey := "latest_weather"
	weatherJSON, err := json.Marshal(weather)
	if err != nil {
		return err
	}
	err = s.cache.SetCache(cacheKey, weatherJSON, 10*time.Minute) // Cache for 10 minutes
	if err != nil {
		return err
	}
	return err
}

func (s *weatherService) GetLatestWeather() (*models.Weather, error) {
	cacheKey := "latest_weather"
	// Check cache first
	cacheWeather, err := s.cache.GetCache(cacheKey)
	if err != nil {
		return nil, err
	}

	if cacheWeather != "" {
		// Return cached data
		var weather models.Weather
		// Assuming you store it as a JSON string, you'd unmarshal it here
		json.Unmarshal([]byte(cacheWeather), &weather)
		// Return the unmarshalled user
		return &weather, nil
	}
	weather, err := s.repo.GetLatestWeather()
	if err != nil {
		return nil, err
	}
	// Cache the result for future requests
	weatherJSON, err := json.Marshal(weather)
	if err != nil {
		return nil, err
	}
	err = s.cache.SetCache(cacheKey, weatherJSON, 10*time.Minute) // Cache for 10 minutes
	if err != nil {
		return nil, err
	}

	return weather, nil
}
