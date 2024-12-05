package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver string
	DBSource string
	Port     string
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env") // Specify the exact filename

	// Add paths where the config file might be located
	viper.AddConfigPath(".")                                                // Current working directory
	viper.AddConfigPath("/home/bunyawat/Documents/my-dev/arch-proj/server") // Explicit path

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	return &Config{
		DBDriver: viper.GetString("DB_DRIVER"),
		DBSource: viper.GetString("DB_SOURCE"),
		Port:     viper.GetString("PORT"),
	}
}
