package impl

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
	"log/slog"
	"net/http"
	"sync"
	"time"
)

type OllamaService struct {
}

var (
	ollamaChatHistory []request_ext.OllamaMessage
	httpClient        *http.Client
	httpClientOnce    sync.Once
)

func NewOllamaService() *OllamaService {
	return &OllamaService{}
}

func (srv *OllamaService) PromptOllamaText(prompt *request.PromptText) (*response_ext.OllamaGenerateResponse, error) {

	jsonRequest, err := prepareOllamaGenerateJsonRequest(prompt)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	url := config.OllamaHost() + "/api/generate"

	ollamaResponse, err := doHttpResponse(ctx, http.MethodPost, url, bytes.NewBuffer((jsonRequest)))
	if err != nil {
		return nil, err
	}

	var response response_ext.OllamaGenerateResponse
	if err := extractHttpResponseIntoValue(ollamaResponse, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}

func (srv *OllamaService) ChatOllama(prompt *request.PromptText) (*response_ext.OllamaChatResponse, error) {
	jsonRequest, err := prepareOllamaChatJsonRequest(prompt)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	url := config.OllamaHost() + "/api/chat"

	ollamaResponse, err := doHttpResponse(ctx, http.MethodPost, url, bytes.NewBuffer((jsonRequest)))
	if err != nil {
		return nil, err
	}

	var response response_ext.OllamaChatResponse
	if err := extractHttpResponseIntoValue(ollamaResponse, &response); err != nil {
		return nil, err
	}

	ollamaChatHistory = append(ollamaChatHistory, request_ext.OllamaMessage{
		Role:    response.Message.Role,
		Content: response.Message.Content,
	})

	return &response, nil
}

func (srv *OllamaService) ResetChat() error {
	ollamaChatHistory = []request_ext.OllamaMessage{}
	return nil
}

func prepareOllamaChatJsonRequest(prompt *request.PromptText) ([]byte, error) {
	ollamaChatHistory = append(ollamaChatHistory, request_ext.OllamaMessage{
		Role:    "user",
		Content: prompt.Text,
	})
	ollamaGenerateRequest := request_ext.OllamaChatRequest{
		Model:    prompt.Model,
		Messages: ollamaChatHistory,
	}

	return json.Marshal(ollamaGenerateRequest)
}

func prepareOllamaGenerateJsonRequest(prompt *request.PromptText) ([]byte, error) {
	ollamaGenerateRequest := request_ext.OllamaGenerateRequest{
		Model:  prompt.Model,
		Prompt: prompt.Text,
	}

	return json.Marshal(ollamaGenerateRequest)
}

func doHttpResponse(ctx context.Context, method string, url string, jsonRequest *bytes.Buffer) (*http.Response, error) {
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

func extractHttpResponseIntoValue[T any](httpResponse *http.Response, response T) error {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("#extractHttpResponseIntoValue - error closing response body", "error", err.Error())
		}
	}(httpResponse.Body)

	if httpResponse.StatusCode != http.StatusOK {
		return extractHttpErrorResponse(httpResponse.Body, httpResponse.StatusCode)
	}

	responseBody, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(responseBody, &response); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil
}

func extractHttpErrorResponse(responseBody io.ReadCloser, statusCode int) error {
	errorResponse, err := io.ReadAll(responseBody)
	if err != nil {
		return fmt.Errorf("API request failed (status: %d): %w", statusCode, err)
	}
	return fmt.Errorf("API request failed (status: %d): %v", statusCode, string(errorResponse))
}
