package config

import (
	"log"
	"os"
)

type Config struct {
	ServerPort   string
	DataFilePath string
	ApiKey       string
}

// LoadConfig loads configuration from environment variables with defaults
func LoadConfig() *Config {
	return &Config{
		ServerPort:   getEnv("SERVER_PORT", "8080"),
		DataFilePath: getEnv("DATA_FILE_PATH", "data/products.json"),
		ApiKey:       getEnv("API_KEY", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Validate if the configuration is valid
func (c *Config) Validate() error {
	if _, err := os.Stat(c.DataFilePath); os.IsNotExist(err) {
		log.Printf("warning: data file path %s does not exist", c.DataFilePath)
	}

	return nil
}
