package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port        string
	Environment string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	RedisHost   string
	RedisPort   string
}

// getEnv returns the value of the environment variable with the given key.
// If the environment variable is not set, it returns an error.
func getEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("environment variable %s is required", key)
	}
	return value, nil
}

// LoadConfig loads the configuration from environment variables.
func LoadConfig() (*Config, error) {
	port, err := getEnv("PORT")
	if err != nil {
		return nil, err
	}

	environment, err := getEnv("ENVIRONMENT")
	if err != nil {
		return nil, err
	}

	DBHost, err := getEnv("DB_HOST")
	if err != nil {
		return nil, err
	}

	DBPort, err := getEnv("DB_PORT")
	if err != nil {
		return nil, err
	}

	DBUser, err := getEnv("DB_USER")
	if err != nil {
		return nil, err
	}

	DBPassword, err := getEnv("DB_PASSWORD")
	if err != nil {
		return nil, err
	}

	DBName, err := getEnv("DB_NAME")
	if err != nil {
		return nil, err
	}

	RedisHost, err := getEnv("REDIS_HOST")
	if err != nil {
		return nil, err
	}

	RedisPort, err := getEnv("REDIS_PORT")
	if err != nil {
		return nil, err
	}

	return &Config{
		Port:        port,
		Environment: environment,
		DBHost:      DBHost,
		DBPort:      DBPort,
		DBUser:      DBUser,
		DBPassword:  DBPassword,
		DBName:      DBName,
		RedisHost:   RedisHost,
		RedisPort:   RedisPort,
	}, nil
}
