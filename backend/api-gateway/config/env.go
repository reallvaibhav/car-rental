package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the API Gateway
type Config struct {
	Server   ServerConfig
	Services ServicesConfig
	JWT      JWTConfig
	CORS     CORSConfig
	Redis    RedisConfig
}

// ServerConfig contains server-related settings
type ServerConfig struct {
	Port  string
	Host  string
	Debug bool
}

// ServicesConfig contains addresses for all microservices
type ServicesConfig struct {
	UserAddr       string
	InventoryAddr  string
	BookingAddr    string
	StatisticsAddr string
}

// JWTConfig contains JWT authentication settings
type JWTConfig struct {
	Secret     string
	Expiration string
}

// CORSConfig contains CORS settings
type CORSConfig struct {
	AllowedOrigins []string
}

// RedisConfig contains Redis connection settings
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	config := &Config{
		Server: ServerConfig{
			Port:  getEnv("PORT", "8080"),
			Host:  getEnv("HOST", "0.0.0.0"),
			Debug: getEnvAsBool("DEBUG", true),
		},
		Services: ServicesConfig{
			UserAddr:       getEnv("USER_SERVICE_ADDR", "localhost:50051"),
			InventoryAddr:  getEnv("INVENTORY_SERVICE_ADDR", "localhost:50052"),
			BookingAddr:    getEnv("BOOKING_SERVICE_ADDR", "localhost:50053"),
			StatisticsAddr: getEnv("STATISTICS_SERVICE_ADDR", "localhost:50054"),
		},
		JWT: JWTConfig{
			Secret:     getEnv("JWT_SECRET", "your-secret-key"),
			Expiration: getEnv("JWT_EXPIRATION", "24h"),
		},
		CORS: CORSConfig{
			AllowedOrigins: getEnvAsSlice("ALLOWED_ORIGINS", []string{"*"}),
		},
		Redis: RedisConfig{
			Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
	}

	return config, nil
}

// Helper functions to get environment variables with default values
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsBool(key string, defaultValue bool) bool {
	strValue := getEnv(key, strconv.FormatBool(defaultValue))
	boolValue, err := strconv.ParseBool(strValue)
	if err != nil {
		log.Printf("Warning: Failed to parse %s as bool: %v", key, err)
		return defaultValue
	}
	return boolValue
}

func getEnvAsInt(key string, defaultValue int) int {
	strValue := getEnv(key, strconv.Itoa(defaultValue))
	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		log.Printf("Warning: Failed to parse %s as int: %v", key, err)
		return defaultValue
	}
	return intValue
}

func getEnvAsSlice(key string, defaultValue []string) []string {
	strValue := getEnv(key, strings.Join(defaultValue, ","))
	if strValue == "" {
		return defaultValue
	}
	return strings.Split(strValue, ",")
}
