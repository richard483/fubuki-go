package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"fubuki-go/config"
	"fubuki-go/dto/request"
	extRequest "fubuki-go/dto/request_ext"
	request2 "fubuki-go/dto/response_ext"
	"github.com/google/generative-ai-go/genai"
	"io"
	"log"
	"net/http"
)

type GeminiApiService struct {
	*genai.Client
}

func NewGeminiApiService(client *genai.Client) *GeminiApiService {
	return &GeminiApiService{client}
}

func (srv *GeminiApiService) PromptText(prompt *request.GeminiText) (error, *[]string) {
	ctx := context.TODO()
	model := srv.Client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt.Text))
	if err != nil {
		return err, nil
	}

	var results []string

	for _, candidate := range resp.Candidates {
		if candidate.Content != nil {
			for _, part := range candidate.Content.Parts {
				if text, ok := part.(genai.Text); ok {
					results = append(results, string(text))
				}
			}
		}
	}

	return nil, &results
}

func (srv *GeminiApiService) Chat(prompt *request.GeminiText) (error, *[]string) {
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent?key=" + config.EnvGeminiApiKey()
	client := &http.Client{}

	var contents []extRequest.GeminiContent
	content := extRequest.GeminiContent{
		Parts: &[]extRequest.GeminiPart{{
			Text: prompt.Text,
		}},
		Role: "user",
	}

	contents = append(contents, content)

	geminiContents := extRequest.GeminiContents{Contents: &contents}

	var jsonData []byte
	jsonData, err := json.Marshal(geminiContents)
	if err != nil {
		return err, nil
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return err, nil
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("API request failed with status code: %d", resp.StatusCode)), nil
	}

	var results []string

	var response request2.GeminiCandidates
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return err, nil
	}

	for _, candidate := range response.Candidates {
		for _, part := range candidate.Content.Parts {
			results = append(results, string(part.Text))
		}
	}

	return nil, &results
}
