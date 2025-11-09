package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("WARN: no .env file using system env")
	}

	cfg := &Config{
		Port: getEnv("PORT", "8080"),
	}

	return cfg
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
