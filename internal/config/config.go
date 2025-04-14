package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	AppEnv            string `mapstructure:"APP_ENV"`
	Port              string `mapstructure:"PORT"`
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DatabaseURL       string `mapstructure:"DATABASE_URL"`
	RedisAddress      string `mapstructure:"REDIS_ADDRESS"`
	RedisDB           int    `mapstructure:"REDIS_DB"`
	RedisPassword     string `mapstructure:"REDIS_PASSWORD"`
	CouponCodeRuneSet string `mapstructure:"COUPON_CODE_RUNESET"`
}

var AppConfig *Config

func (c *Config) IsProduction() bool {
	return c.AppEnv == "production"
}

func Init() {
	_ = godotenv.Load()

	redisDB := getEnv("REDIS_DB", "0")
	redisDBInt, err := strconv.Atoi(redisDB)
	if err != nil {
		redisDBInt = 0
	}

	AppConfig = &Config{
		AppEnv:            getEnv("APP_ENV", "development"),
		Port:              getEnv("PORT", "8000"),
		DBDriver:          getEnv("DB_DRIVER", "sqlite"),
		DatabaseURL:       getEnv("DATABASE_URL", ""),
		RedisAddress:      getEnv("REDIS_ADDRESS", ""),
		RedisDB:           redisDBInt,
		RedisPassword:     getEnv("REDIS_PASSWORD", ""),
		CouponCodeRuneSet: getEnv("COUPON_CODE_RUNESET", "0123456789가나다라마바사아자차카타파하"),
	}
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}
