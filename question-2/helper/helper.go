package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetEnv function to get .env file
func GetEnv(key string) string {
	// load .env file
	switch godotenv.Load() {
	case godotenv.Load("../.env"):
		log.Println("Error loading .env file")
	case godotenv.Load("../../.env"):
		log.Println("Error loading .env file")
	}
	return os.Getenv(key)
}
