package config

import (
	"fmt"
	"os"
)

type Config struct {
    APIKey string
}

func NewConfig() (*Config, error) {
    apiKey := os.Getenv("WEATHER_API_KEY")
    if apiKey == "" {
        return nil, fmt.Errorf("WEATHER_API_KEY environment variable is not set")
    }
    return &Config{APIKey: apiKey}, nil
}
