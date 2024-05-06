package config

import (
	"log"
	"os"
	"strconv"
)

func EnvGeminiApiKey() string {
	return os.Getenv("GEMINI_API_KEY")
}

func EnvPostgresURI() string {
	return os.Getenv("POSTGRES_URI")
}

func EnvGeminiAPI() bool {
	res, err := strconv.ParseBool(os.Getenv("GEMINI_API"))
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return res
}
