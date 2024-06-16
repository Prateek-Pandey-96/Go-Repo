package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ConnectionURI  string
	DatabaseName   string
	Port           string
	CollectionName string
}

var AppConfig Config

func LoadConfig() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	AppConfig = Config{
		ConnectionURI:  os.Getenv("CONNECTION_URI"),
		DatabaseName:   os.Getenv("DATABASE"),
		Port:           os.Getenv("PORT"),
		CollectionName: os.Getenv("COLLECTION"),
	}

	if AppConfig.ConnectionURI == "" ||
		AppConfig.DatabaseName == "" ||
		AppConfig.Port == "" ||
		AppConfig.CollectionName == "" {
		log.Fatal("Missing required environment variables")
	}
}
