package config

import (
	"log/slog"
	"os"
	"strconv"
)

func EnvGeminiApiKey() string {
	return os.Getenv("GEMINI_API_KEY")
}

func EnvPostgresURI() string {
	value, present := os.LookupEnv("POSTGRES_URI")
	if !present {
		slog.Warn("#EnvPostgresURI - POSTGRES_URI not set, using default value")
		return "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	}
	return value
}

func EnvRedisURI() string {
	value, present := os.LookupEnv("REDIS_URI")
	if !present {
		slog.Warn("#EnvRedisURI - REDIS_URI not set, using default value")
		return "redis://localhost:6379/0"
	}
	return value
}

func EnvHost() string {
	value, present := os.LookupEnv("HOST")
	if !present {
		slog.Warn("#EnvHost - HOST not set, using default value")
		return "localhost"
	}
	return value
}

func EnvRetrieveHistory() bool {
	res, err := strconv.ParseBool(os.Getenv("RETRIEVE_HISTORY"))
	if err != nil {
		slog.Error("#EnvRetrieveHistory - error parsing RETRIEVE_HISTORY", "error", err.Error())
		return false
	}
	return res
}

func EnvReleaseMode() bool {
	res, err := strconv.ParseBool(os.Getenv("RELEASE_MODE"))
	if err != nil {
		slog.Error("#EnvReleaseMode - error parsing RELEASE_MODE", "error", err.Error())
		return false
	}
	return res
}

func EnvPort() string {
	value, present := os.LookupEnv("PORT")
	if !present {
		slog.Warn("#EnvPort - PORT not set, using default value")
		return "8080"
	}
	return value
}

func EnvGeminiModel() string {
	value, present := os.LookupEnv("GEMINI_MODEL")
	if !present {
		slog.Warn("#EnvGeminiModel - GEMINI_MODEL not set, using default value")
		return "gemini-1.5-turbo"
	}
	return value
}

func OllamaHost() string {
	value, present := os.LookupEnv("OLLAMA_HOST")
	if !present {
		slog.Warn("#OllamaHost - OLLAMA_HOST not set, using default value")
		return "http://localhost:11434"

	}
	return value
}
