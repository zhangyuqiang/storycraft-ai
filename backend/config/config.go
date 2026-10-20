package config

import (
	"os"
)

type Config struct {
	Port            string
	LLMBaseURL      string
	LLMAPIKey       string
	SeedanceBaseURL string
	SeedanceAPIKey  string
}

func Load() Config {
	return Config{
		Port:            getEnv("PORT", "8080"),
		LLMBaseURL:      getEnv("LLM_BASE_URL", "http://localhost:11434"),
		LLMAPIKey:       getEnv("LLM_API_KEY", "demo-key"),
		SeedanceBaseURL: getEnv("SEEDANCE_BASE_URL", "https://api.seedance.example.com"),
		SeedanceAPIKey:  getEnv("SEEDANCE_API_KEY", "demo-key"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
