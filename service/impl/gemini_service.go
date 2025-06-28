package impl

import (
	"context"
	"fubuki-go/config"
	"fubuki-go/dto/request"
	"fubuki-go/repository"

	genai "google.golang.org/genai"
)

type GeminiService struct {
	*genai.Client
	repository.HistoryRepositoryInterface
	repository.CacheRepositoryInterface
}

// NewGeminiService creates and returns a new GeminiService with the provided AI client, history repository, and cache repository.
func NewGeminiService(client *genai.Client, repository repository.HistoryRepositoryInterface, cache repository.CacheRepositoryInterface) *GeminiService {
	return &GeminiService{client, repository, cache}
}

var geminiContent []*genai.Content

func (srv *GeminiService) ResetSession() (string, error) {
	geminiContent = nil
	return "ok", nil
}

func (srv *GeminiService) PromptText(prompt *request.PromptText) (string, error) {
	ctx := context.TODO()
	resp, err := srv.Client.Models.GenerateContent(ctx, config.EnvGeminiModel(), genai.Text(prompt.Text), srv.getGenerateContentConfig())
	if err != nil {
		return "", err
	}

	return resp.Text(), nil
}

func (srv *GeminiService) Chat(prompt *request.PromptText) (string, error) {
	ctx := context.TODO()

	if config.EnvRetrieveHistory() {
		var histories = srv.GetAllByModelSource("gemini")

		for _, history := range histories {
			geminiContent = append(geminiContent, genai.NewContentFromText(history.UserQuestion, genai.RoleUser))
			geminiContent = append(geminiContent, genai.NewContentFromText(history.ModelAnswer, genai.RoleModel))
		}
	}

	chat, _ := srv.Client.Chats.Create(ctx, config.EnvGeminiModel(), srv.getGenerateContentConfig(), geminiContent)
	res, err := chat.SendMessage(ctx, genai.Part{Text: prompt.Text})

	if err != nil {
		return "", err
	}

	geminiContent = append(geminiContent, genai.NewContentFromText(prompt.Text, genai.RoleUser))
	geminiContent = append(geminiContent, genai.NewContentFromText(res.Text(), genai.RoleModel))

	return res.Text(), nil
}

func (srv *GeminiService) getGenerateContentConfig() *genai.GenerateContentConfig {
	if config.EnvGeminiGoogleSearch() {
		return &genai.GenerateContentConfig{
			Tools: []*genai.Tool{
				{GoogleSearch: &genai.GoogleSearch{}},
			},
		}
	} else {
		return nil
	}
}
