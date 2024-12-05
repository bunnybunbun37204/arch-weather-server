package handlers

import (
	"net/http"
	"server/internal/models"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

type WeatherHandler struct {
	service services.WeatherService
}

func NewWetherHandler(service services.WeatherService) *WeatherHandler {
	return &WeatherHandler{service: service}
}

func (h *WeatherHandler) CreateWeather(c *gin.Context) {
	var weather models.Weather
	if err := c.ShouldBindJSON(&weather); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateWeather(&weather); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, weather)
}

func (h *WeatherHandler) GetLatestWeather(c *gin.Context) {

	weather, err := h.service.GetLatestWeather()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Weather not found"})
		return
	}

	c.JSON(http.StatusOK, weather)
}
