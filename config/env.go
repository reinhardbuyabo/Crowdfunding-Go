// config/env.go
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	// Database Configuration
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// Server Configuration
	ServerPort string

	// Blockchain Configuration
	BlockchainNodeURL string
	ContractAddress   string

	// Logging Configuration
	LogLevel string
}

// LoadConfig reads configuration from environment variables or .env file
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		// It's okay if no .env file exists
		fmt.Println("No .env file found. Using system environment variables.")
	}

	config := &Config{
		// Database Configuration
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "crowdfunding"),

		// Server Configuration
		ServerPort: getEnv("SERVER_PORT", "8080"),

		// Blockchain Configuration
		BlockchainNodeURL: getEnv("BLOCKCHAIN_NODE_URL", "https://mainnet.infura.io/v3/"),
		ContractAddress:   getEnv("CONTRACT_ADDRESS", ""),

		// Logging Configuration
		LogLevel: getEnv("LOG_LEVEL", "info"),
	}

	// Validate required configurations
	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

// getEnv retrieves an environment variable with a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// validateConfig performs basic validation of configuration
func validateConfig(config *Config) error {
	// Add specific validation rules
	if config.DBHost == "" {
		return fmt.Errorf("DB_HOST is required")
	}

	if config.DBPassword == "" {
		return fmt.Errorf("DB_PASSWORD is required")
	}

	return nil
}