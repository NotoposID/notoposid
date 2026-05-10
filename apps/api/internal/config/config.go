package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv    string
	Port      string
	DBURL     string
	RedisURL  string
	JWTSecret string
	OpenAIKey string
}

func LoadConfig() *Config {
	godotenv.Load()

	return &Config{
		AppEnv:    getEnv("APP_ENV", "development"),
		Port:      getEnv("PORT", "8000"),
		DBURL:     getEnv("DATABASE_URL", ""),
		RedisURL:  getEnv("REDIS_URL", ""),
		JWTSecret: getEnv("JWT_SECRET", "supersecret"),
		OpenAIKey: getEnv("OPENAI_API_KEY", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
