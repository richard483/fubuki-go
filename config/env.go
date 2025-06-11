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
	value, present := os.LookupEnv("POSTGRES_URI")
	if !present {
		log.Println("#ERROR POSTGRES_URI not set, using default value")
		return "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	}
	return value
}

func EnvRedisURI() string {
	return os.Getenv("REDIS_URI")
}

func EnvHost() string {
	value, present := os.LookupEnv("HOST")
	if !present {
		log.Println("#ERROR HOST not set, using default value")
		return "localhost"
	}
	return value
}

func EnvRetrieveHistory() bool {
	res, err := strconv.ParseBool(os.Getenv("RETRIEVE_HISTORY"))
	if err != nil {
		log.Println("#ERROR " + err.Error())
		return false
	}
	return res
}

func EnvReleaseMode() bool {
	res, err := strconv.ParseBool(os.Getenv("RELEASE_MODE"))
	if err != nil {
		log.Println("#ERROR " + err.Error())
		return false
	}
	return res
}

func EnvPort() string {
	value, present := os.LookupEnv("PORT")
	if !present {
		log.Println("#ERROR PORT not set, using default value")
		return "8080"
	}
	return value
}

func EnvGeminiModel() string {
	value, present := os.LookupEnv("GEMINI_MODEL")
	if !present {
		log.Println("#ERROR GEMINI_MODEL not set, using default value")
		return "gemini-1.5-turbo"
	}
	return value
}

func OllamaHost() string {
	value, present := os.LookupEnv("OLLAMA_HOST")
	if !present {
		log.Println("#ERROR OLLAMA_HOST not set, using default value")
		return "http://localhost:11434"

	}
	return value
}
