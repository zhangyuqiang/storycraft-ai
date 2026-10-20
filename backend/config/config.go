package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port            string
	LLMBaseURL      string
	LLMAPIKey       string
	SeedanceBaseURL string
	SeedanceAPIKey  string
	DBHost          string
	DBPort          string
	DBUser          string
	DBPassword      string
	DBName          string
	DBCharset       string
	DBTimezone      string
}

func Load() Config {
	return Config{
		Port:            getEnv("PORT", "8080"),
		LLMBaseURL:      getEnv("LLM_BASE_URL", "http://localhost:11434"),
		LLMAPIKey:       getEnv("LLM_API_KEY", "demo-key"),
		SeedanceBaseURL: getEnv("SEEDANCE_BASE_URL", "https://api.seedance.example.com"),
		SeedanceAPIKey:  getEnv("SEEDANCE_API_KEY", "demo-key"),
		DBHost:          getEnv("DB_HOST", "127.0.0.1"),
		DBPort:          getEnv("DB_PORT", "3306"),
		DBUser:          getEnv("DB_USER", "root"),
		DBPassword:      getEnv("DB_PASSWORD", "123456"),
		DBName:          getEnv("DB_NAME", "storycraft_ai"),
		DBCharset:       getEnv("DB_CHARSET", "utf8mb4"),
		DBTimezone:      getEnv("DB_TIMEZONE", "Local"),
	}
}

func (c Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBCharset,
		c.DBTimezone,
	)
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
