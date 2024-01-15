package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var envVars map[string]string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env file couldn't be loaded")
	}

	envVars = make(map[string]string)
	loadEnvVariable("JWT_SECRET_KEY")
	// Add more variables here if needed
	// loadEnvVariable("API_SECRET")
	// Add more variables here if needed
}

func loadEnvVariable(key string) {
	value := os.Getenv(key)
	envVars[key] = value
}

func GetEnvVariable(key string) string {
	return envVars[key]
}
