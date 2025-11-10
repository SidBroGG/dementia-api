package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DB struct {
	Host     string
	Name     string
	User     string
	Password string
	Port     string
}

type Config struct {
	DB

	Port string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("WARN: no .env file using system env")
	}

	cfg := &Config{}

	cfg.Port = getEnv("PORT", "8080")
	cfg.DB.Host = getEnv("DB_HOST", "db")
	cfg.DB.Name = getEnv("DB_NAME", "todolist_db")
	cfg.DB.User = getEnv("DB_USER", "user")
	cfg.DB.Password = getEnv("DB_PASSWORD", "password")
	cfg.DB.Port = getEnv("DB_PORT", "5432")

	return cfg
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
