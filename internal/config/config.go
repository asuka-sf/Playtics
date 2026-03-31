package config

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func Load() *Config {
	// load the config from the environment variables
	_ = godotenv.Load()

	cfg := &Config{
		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", ""),
	}

	if cfg.DBUser == "" {
		log.Fatal("DB_USER environment variable is required")
	}
	if cfg.DBName == "" {
		log.Fatal("DB_NAME environment variable is required")
	}

	return cfg
}

// return the value of the environment variable if it exists
func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// DatabaseURL assembles a connection string from individual config fields
func (c *Config) DatabaseURL() string {
	u := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(c.DBUser, c.DBPassword),
		Host:     fmt.Sprintf("%s:%s", c.DBHost, c.DBPort),
		Path:     c.DBName,
		RawQuery: "sslmode=disable",
	}
	return u.String()
}
