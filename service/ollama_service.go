package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	config "fubuki-go/config"
	"fubuki-go/dto/request"
	"fubuki-go/dto/request_ext"
	"fubuki-go/dto/response_ext"
	repository "fubuki-go/repository_interface"
	"io"
	"log"
	"net/http"

	"github.com/google/generative-ai-go/genai"
)

type OllamaService struct {
	*genai.Client
	repository.GeminiHistoryRepositoryInterface
}

func NewOllamaService(client *genai.Client, repository repository.GeminiHistoryRepositoryInterface) *GeminiService {
	return &GeminiService{client, repository}
}

func (srv *GeminiService) PromptOllamaText(prompt *request.PromptText) (*response_ext.OllamaGenerateResponse, error) {
	url := config.OllamaHost() + "/api/generate"
	ollamaGenerateRequest := request_ext.OllamaGenerateRequest{
		Model:  prompt.Model,
		Prompt: prompt.Text,
		Stream: "false",
	}

	var jsonRequest []byte
	jsonRequest, err := json.Marshal(ollamaGenerateRequest)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer((jsonRequest)))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	ollamaResponse, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("#ERROR " + err.Error())
		}
	}(ollamaResponse.Body)

	if ollamaResponse.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("API request failed with status code: %d", ollamaResponse.StatusCode))
	}

	responseBody, err := io.ReadAll(ollamaResponse.Body)
	if err != nil {
		return nil, err
	}

	var response response_ext.OllamaGenerateResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		log.Println("#ERROR " + err.Error())
		return nil, err
	}

	return &response, nil
}
