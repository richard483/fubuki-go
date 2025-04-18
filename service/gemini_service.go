package service

import (
	"context"
	"fubuki-go/config"
	"fubuki-go/dto/request"
	repository "fubuki-go/repository_interface"
	"github.com/google/generative-ai-go/genai"
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

func (srv *GeminiService) geminiModel() *genai.GenerativeModel {
	if geminiModel == nil {
		geminiModel = srv.Client.GenerativeModel("gemini-2.0-flash")
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
