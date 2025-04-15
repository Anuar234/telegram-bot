package config

import (
	"os"
)

// Config holds all configuration for the bot
type Config struct {
	TelegramToken string
	Debug         bool
}

// Load loads configuration from environment or defaults
func Load() *Config {
	token := os.Getenv("7761997881:AAEH98PKYzxFFpFZdP1Xzn8JAM76RPI0cPY")
	if token == "" {
		token = "7761997881:AAEH98PKYzxFFpFZdP1Xzn8JAM76RPI0cPY" // Default, but better to use environment variable
	}
	
	return &Config{
		TelegramToken: token,
		Debug:         true,
	}
}