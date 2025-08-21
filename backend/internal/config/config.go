package config

import (
	"os"
	"strconv"
	"strings"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Logging  LoggingConfig
	CORS     CORSConfig
	Env      string
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port string
	Host string
}

// DatabaseConfig holds database-related configuration
type DatabaseConfig struct {
	Path string
}

// LoggingConfig holds logging-related configuration
type LoggingConfig struct {
	Requests bool
	SQL      bool
	CORS     bool
}

// CORSConfig holds CORS-related configuration
type CORSConfig struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
	MaxAge           int
}

// Load loads configuration from environment variables
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("BACKEND_PORT", "6060"),
			Host: getEnv("BACKEND_HOST", "localhost"),
		},
		Database: DatabaseConfig{
			Path: getEnv("DB_PATH", "finance.db"),
		},
		Logging: LoggingConfig{
			Requests: getEnvBool("LOG_REQUESTS", true),
			SQL:      getEnvBool("LOG_SQL", false),
			CORS:     getEnvBool("LOG_CORS", false),
		},
		CORS: CORSConfig{
			AllowOrigins:     strings.Split(getEnv("CORS_ALLOW_ORIGINS", "*"), ","),
			AllowMethods:     strings.Split(getEnv("CORS_ALLOW_METHODS", "GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS"), ","),
			AllowHeaders:     strings.Split(getEnv("CORS_ALLOW_HEADERS", "Origin,Content-Length,Content-Type,Accept,Authorization,X-Requested-With"), ","),
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: false,
			MaxAge:           getEnvInt("CORS_MAX_AGE", 43200),
		},
		Env: getEnv("ENV", "development"),
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvBool gets a boolean environment variable or returns a default value
func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.ParseBool(value); err == nil {
			return parsed
		}
	}
	return defaultValue
}

// getEnvInt gets an integer environment variable or returns a default value
func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil {
			return parsed
		}
	}
	return defaultValue
}
