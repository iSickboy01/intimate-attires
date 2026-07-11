package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	Database_URL string
}

func LoadConfig() *Config {

	err := godotenv.Load()

	if err != nil {
		log.Printf("env file is missing dawg!:%s", err)
	}

	cfg := &Config{
		Port:         os.Getenv("Port"),
		Database_URL: os.Getenv("Database_URL"),
	}

	return cfg
}
