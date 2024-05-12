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
	repository "fubuki-go/repository_interface"
	"github.com/google/generative-ai-go/genai"
	"io"
	"log"
	"net/http"
)

type GeminiService struct {
	*genai.Client
	repository.GeminiHistoryRepositoryInterface
}

func NewGeminiService(client *genai.Client, repository repository.GeminiHistoryRepositoryInterface) *GeminiService {
	return &GeminiService{client, repository}
}

var geminiModel *genai.GenerativeModel
var chatSession *genai.ChatSession

func (srv *GeminiService) ResetSession() (error, string) {
	geminiModel = nil
	chatSession = nil
	return nil, "ok"
}

func (srv *GeminiService) PromptText(prompt *request.GeminiText) (error, *[]string) {
	ctx := context.TODO()
	model := srv.geminiModel()
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

func (srv *GeminiService) Chat(prompt *request.GeminiText) (error, *[]string) {
	ctx := context.TODO()

	model := srv.geminiModel()
	cs := srv.chatSession(model)

	if config.EnvRetrieveHistory() {
		var histories = srv.GetAll()

		for _, history := range histories {
			cs.History = append(cs.History, &genai.Content{
				Parts: []genai.Part{
					genai.Text(history.UserQuestion),
				},
				Role: "user",
			})

			cs.History = append(cs.History, &genai.Content{
				Parts: []genai.Part{
					genai.Text(history.ModelAnswer),
				},
				Role: "model",
			})
		}
	}

	resp, err := cs.SendMessage(ctx, genai.Text(prompt.Text))

	if err != nil {
		cs.History = cs.History[:len(cs.History)-1]
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

func (srv *GeminiService) TuneModel() (error, *request2.GeminiTunedModel) {
	url := "https://generativelanguage.googleapis.com/v1beta/tunedModels"
	client := &http.Client{}

	var geminiExamplesTrainingData []extRequest.GeminiExampleTrainingData

	var histories = srv.GetAll()
	for _, history := range histories {
		geminiExamplesTrainingData = append(geminiExamplesTrainingData, extRequest.GeminiExampleTrainingData{
			TextInput: history.UserQuestion,
			Output:    history.ModelAnswer,
		})
	}

	hyperparameter := extRequest.GeminiHyperparameter{
		BatchSize:    2,
		LearningRate: 0.001,
		EpochCount:   5,
	}

	examplesTrainingData := extRequest.GeminiExamplesTrainingData{Examples: &geminiExamplesTrainingData}
	trainingData := extRequest.GeminiTrainingData{Examples: &examplesTrainingData}

	tuningTask := extRequest.GeminiTuningTask{
		Hyperparameter: &hyperparameter,
		TrainingData:   &trainingData,
	}

	tuneModel := extRequest.GeminiTuneModel{
		BaseModel:   "models/gemini-1.0-pro-001",
		DisplayName: "FBK base model",
		TuningTask:  &tuningTask,
	}

	var jsonData []byte
	jsonData, err := json.Marshal(tuneModel)
	if err != nil {
		return err, nil
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonData))
	if err != nil {
		return err, nil
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", config.EnvGoogleAccessToken())
	req.Header.Set("x-goog-user-project", config.EnvGoogleProjectId())
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

	var response request2.GeminiTunedModel
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return err, nil
	}

	return nil, &response
}

func (srv *GeminiService) geminiModel() *genai.GenerativeModel {
	if geminiModel == nil {
		geminiModel = srv.Client.GenerativeModel("gemini-pro")
		geminiModel.SafetySettings = []*genai.SafetySetting{
			{
				Category:  genai.HarmCategoryHarassment,
				Threshold: genai.HarmBlockNone,
			},
			{
				Category:  genai.HarmCategoryHateSpeech,
				Threshold: genai.HarmBlockNone,
			},
			{
				Category:  genai.HarmCategoryDangerousContent,
				Threshold: genai.HarmBlockNone,
			},
			{
				Category:  genai.HarmCategorySexuallyExplicit,
				Threshold: genai.HarmBlockNone,
			},
		}
	}
	return geminiModel
}

func (srv *GeminiService) chatSession(model *genai.GenerativeModel) *genai.ChatSession {
	if chatSession == nil {
		chatSession = model.StartChat()
	}
	return chatSession
}
