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

func EnvHost() string {
	return os.Getenv("HOST")
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
	return os.Getenv("PORT")
}

func EnvGeminiModel() string {
	return os.Getenv("GEMINI_MODEL")
}
