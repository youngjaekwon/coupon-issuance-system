package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	AppEnv       string `mapstructure:"APP_ENV"`
	Port         string `mapstructure:"PORT"`
	DBDriver     string `mapstructure:"DB_DRIVER"`
	DatabaseURL  string `mapstructure:"DATABASE_URL"`
	RedisAddress string `mapstructure:"REDIS_ADDRESS"`
}

var AppConfig *Config

func (c *Config) IsProduction() bool {
	return c.AppEnv == "production"
}

func Init() {
	_ = godotenv.Load()

	AppConfig = &Config{
		AppEnv:       getEnv("APP_ENV", "development"),
		Port:         getEnv("PORT", "8000"),
		DBDriver:     getEnv("DB_DRIVER", "sqlite"),
		DatabaseURL:  getEnv("DATABASE_URL", ""),
		RedisAddress: getEnv("REDIS_ADDRESS", ""),
	}
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}
