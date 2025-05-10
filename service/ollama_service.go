package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	config "fubuki-go/config"
	"fubuki-go/dto/request"
	"fubuki-go/dto/request_ext"
	"fubuki-go/dto/response_ext"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type OllamaService struct {
}

var (
	httpClient     *http.Client
	httpClientOnce sync.Once
)

func NewOllamaService() *OllamaService {
	return &OllamaService{}
}

func (srv *OllamaService) PromptOllamaText(prompt *request.PromptText) (*response_ext.OllamaGenerateResponse, error) {

	jsonRequest, err := prepareJsonRequest(prompt)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	url := config.OllamaHost() + "/api/generate"

	ollamaResponse, err := getHttpResponse(ctx, http.MethodPost, url, bytes.NewBuffer((jsonRequest)))
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
		return nil, extractOllamaErrorResponse(ollamaResponse.Body, ollamaResponse.StatusCode)
	}

	responseBody, err := io.ReadAll(ollamaResponse.Body)
	if err != nil {
		return nil, err
	}

	var response response_ext.OllamaGenerateResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}

func prepareJsonRequest(prompt *request.PromptText) ([]byte, error) {
	ollamaGenerateRequest := request_ext.OllamaGenerateRequest{
		Model:  prompt.Model,
		Prompt: prompt.Text,
		Stream: false,
	}

	return json.Marshal(ollamaGenerateRequest)
}

func getHttpResponse(ctx context.Context, method string, url string, jsonRequest *bytes.Buffer) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, jsonRequest)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return getHttpClient().Do(req)
}

func getHttpClient() *http.Client {
	httpClientOnce.Do(func() {
		httpClient = &http.Client{
			Timeout: 100 * time.Second,
		}
	})
	return httpClient
}

func extractOllamaErrorResponse(responseBody io.ReadCloser, statusCode int) error {
	errorResponse, err := io.ReadAll(responseBody)
	if err != nil {
		return fmt.Errorf("API request failed (status: %d): %w", statusCode, err)
	}
	return fmt.Errorf("API request failed (status: %d): %w", statusCode, string(errorResponse))
}
