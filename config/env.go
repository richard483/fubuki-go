package config

import (
	"os"
)

func EnvGeminiApiKey() string {
	return os.Getenv("GEMINI_API_KEY")
}

func EnvPostgresURI() string {
	return os.Getenv("POSTGRES_URI")
}
