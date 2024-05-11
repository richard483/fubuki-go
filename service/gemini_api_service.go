package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"fubuki-go/config"
	"fubuki-go/dto/request"
	extRequest "fubuki-go/dto/request_ext"
	request2 "fubuki-go/dto/response_ext"
	repository "fubuki-go/repository_interface"
	"github.com/google/generative-ai-go/genai"
	"io"
	"log"
	"net/http"
	"sync"
)

type GeminiApiService struct {
	*genai.Client
	repository.GeminiHistoryRepositoryInterface
}

var lock = &sync.Mutex{}
var historyContent []extRequest.GeminiContent
var geminiDefaultService *GeminiService

func NewGeminiApiService(client *genai.Client, repository repository.GeminiHistoryRepositoryInterface) *GeminiApiService {
	return &GeminiApiService{client, repository}
}

func (srv *GeminiApiService) PromptText(prompt *request.GeminiText) (error, *[]string) {
	if geminiDefaultService == nil {
		geminiDefaultService = NewGeminiService(srv.Client, srv.GeminiHistoryRepositoryInterface)
	}
	return geminiDefaultService.PromptText(prompt)
}

func (srv *GeminiApiService) Chat(prompt *request.GeminiText) (error, *[]string) {
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent?key=" + config.EnvGeminiApiKey()
	client := &http.Client{}

	if historyContent == nil && config.EnvRetrieveHistory() {
		lock.Lock()
		defer lock.Unlock()

		if historyContent == nil {
			var histories = srv.GetAll()
			for _, history := range histories {
				historyContent = append(historyContent, extRequest.GeminiContent{
					Parts: &[]extRequest.GeminiPart{{
						Text: history.UserQuestion,
					}},
					Role: "user",
				})

				historyContent = append(historyContent, extRequest.GeminiContent{
					Parts: &[]extRequest.GeminiPart{{
						Text: history.ModelAnswer,
					}},
					Role: "model",
				})
			}
		}
	}

	content := extRequest.GeminiContent{
		Parts: &[]extRequest.GeminiPart{{
			Text: prompt.Text,
		}},
		Role: "user",
	}

	historyContent = append(historyContent, content)

	geminiContents := extRequest.GeminiContents{Contents: &historyContent}

	var jsonData []byte
	jsonData, err := json.Marshal(geminiContents)
	if err != nil {
		return err, nil
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonData))
	if err != nil {
		return err, nil
	}

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
	if len(results) == 0 {
		return errors.New(fmt.Sprintf("Zero response")), nil
	}

	content = extRequest.GeminiContent{
		Parts: &[]extRequest.GeminiPart{{
			Text: results[0],
		}},
		Role: "model",
	}

	if len(historyContent) > 500 {
		historyContent = append(historyContent[:0], historyContent[1:]...)
	}
	historyContent = append(historyContent, content)

	return nil, &results
}

func (srv *GeminiApiService) TuneModel() (error, *request2.GeminiTunedModel) {
	if geminiDefaultService == nil {
		geminiDefaultService = NewGeminiService(srv.Client, srv.GeminiHistoryRepositoryInterface)
	}
	return geminiDefaultService.TuneModel()
}
