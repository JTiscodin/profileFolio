// config/config.go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func GetMongoURI() string {
	uri, exists := os.LookupEnv("MONGODB_URI")
	if !exists {
		log.Fatal("MONGODB_URI not set in environment")
	}
	return uri
}


func GetFrontendURI() string {
	uri, exists := os.LookupEnv("ALLOWED_ORIGINS")
	if !exists {
		log.Fatal("ALLOWED_ORIGINS not set in environment")
	}
	log.Print("FrontendURI is ", uri)
	return uri
}

func GetOpenAIAPIKey() string {
	uri, exists := os.LookupEnv("OPENAI_API_KEY")
	if !exists {
		log.Fatal("OPENAI_API_KEY not set in environment")
	}
	return uri
}
