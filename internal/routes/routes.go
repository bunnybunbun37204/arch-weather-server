package routes

import (
	"server/internal/cache"
	"server/internal/handlers"
	"server/internal/repositories"
	"server/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB, cache *cache.RedisCache) {
	weatherRepo := repositories.NewWeatherRepository(db)
	weatherService := services.NewWeatherService(weatherRepo, cache)
	weatherHandler := handlers.NewWetherHandler(weatherService)

	weatherRoute := r.Group("/api2/weather")
	{
		weatherRoute.POST("", weatherHandler.CreateWeather)
		weatherRoute.GET("", weatherHandler.GetLatestWeather)
	}
}
