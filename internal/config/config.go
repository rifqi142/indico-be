package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName          string
	AppEnv           string
	AppPort          string
	DBHost           string
	DBPort           string
	DBUser           string
	DBPassword       string
	DBName           string
	DBSSLMode        string
	JWTSecret        string
	JWTExpiration    string
	ServerReadTimeout  string
	ServerWriteTimeout string
}

func LoadConfig() *Config {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config := &Config{
		AppName:            getEnv("APP_NAME", "indico-be"),
		AppEnv:             getEnv("APP_ENV", "development"),
		AppPort:            getEnv("APP_PORT", "8080"),
		DBHost:             getEnv("DB_HOST", "localhost"),
		DBPort:             getEnv("DB_PORT", "5432"),
		DBUser:             getEnv("DB_USER", "postgres"),
		DBPassword:         getEnv("DB_PASSWORD", "rajawali02"),
		DBName:             getEnv("DB_NAME", "indico_db"),
		DBSSLMode:          getEnv("DB_SSL_MODE", "disable"),
		JWTSecret:          getEnv("JWT_SECRET", "your_secret_key"),
		JWTExpiration:      getEnv("JWT_EXPIRATION", "24h"),
		ServerReadTimeout:  getEnv("SERVER_READ_TIMEOUT", "10s"),
		ServerWriteTimeout: getEnv("SERVER_WRITE_TIMEOUT", "10s"),
	}

	return config
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode,
	)
}
