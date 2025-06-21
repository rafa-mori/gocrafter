package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	Environment  string
	Port         string
	LogLevel     string
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int

	{{- if .DatabaseType}}
	// Database configuration
	DatabaseURL      string
	DatabaseHost     string
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabaseSSLMode  string
	{{- end}}

	{{- if .CacheType}}
	// Cache configuration
	{{- if eq .CacheType "redis"}}
	RedisURL      string
	RedisPassword string
	RedisDB       int
	{{- end}}
	{{- end}}

	{{- if hasFeature "Authentication"}}
	// JWT configuration
	JWTSecret     string
	JWTExpiration int
	{{- end}}
}

// Load loads the configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	config := &Config{
		Environment:  getEnv("ENVIRONMENT", "development"),
		Port:         getEnv("PORT", "8080"),
		LogLevel:     getEnv("LOG_LEVEL", "info"),
		ReadTimeout:  getEnvAsInt("READ_TIMEOUT", 10),
		WriteTimeout: getEnvAsInt("WRITE_TIMEOUT", 10),
		IdleTimeout:  getEnvAsInt("IDLE_TIMEOUT", 60),

		{{- if .DatabaseType}}
		// Database
		DatabaseURL:      getEnv("DATABASE_URL", ""),
		DatabaseHost:     getEnv("DB_HOST", "localhost"),
		DatabasePort:     getEnv("DB_PORT", "{{- if eq .DatabaseType "postgres"}}5432{{- else if eq .DatabaseType "mysql"}}3306{{- else if eq .DatabaseType "mongodb"}}27017{{- else}}3306{{- end}}"),
		DatabaseUser:     getEnv("DB_USER", "{{- if eq .DatabaseType "postgres"}}postgres{{- else}}root{{- end}}"),
		DatabasePassword: getEnv("DB_PASSWORD", ""),
		DatabaseName:     getEnv("DB_NAME", "{{.PackageName}}"),
		DatabaseSSLMode:  getEnv("DB_SSL_MODE", "disable"),
		{{- end}}

		{{- if .CacheType}}
		{{- if eq .CacheType "redis"}}
		// Redis
		RedisURL:      getEnv("REDIS_URL", "redis://localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvAsInt("REDIS_DB", 0),
		{{- end}}
		{{- end}}

		{{- if hasFeature "Authentication"}}
		// JWT
		JWTSecret:     getEnv("JWT_SECRET", "your-secret-key-change-this-in-production"),
		JWTExpiration: getEnvAsInt("JWT_EXPIRATION", 3600), // 1 hour
		{{- end}}
	}

	return config, nil
}

// getEnv gets an environment variable with a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt gets an environment variable as integer with a default value
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
