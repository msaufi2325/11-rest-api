package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnv(v string) (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(".env file couldn't be loaded")
	}

	key := os.Getenv(v)

	return key, nil
}
