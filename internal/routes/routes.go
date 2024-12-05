package routes

import (
	"server/internal/handlers"
	"server/internal/repositories"
	"server/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	weatherRepo := repositories.NewWeatherRepository(db)
	weatherService := services.NewWeatherService(weatherRepo)
	weatherHandler := handlers.NewWetherHandler(weatherService)

	weatherRoute := r.Group("/weather")
	{
		weatherRoute.POST("", weatherHandler.CreateWeather)
		weatherRoute.GET("", weatherHandler.GetLatestWeather)
	}
}
