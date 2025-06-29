package impl

import (
	"context"
	"errors"
	"fubuki-go/config"
	"fubuki-go/dto/request"
	"fubuki-go/helper"
	"fubuki-go/repository"
	"sync"

	"github.com/redis/go-redis/v9"
	genai "google.golang.org/genai"
)

type GeminiService struct {
	*genai.Client
	repository.HistoryRepositoryInterface
	cache repository.CacheRepositoryInterface
}

var (
	geminiRedisCacheKey string
	geminiServiceOnce   sync.Once
)

func NewGeminiService(client *genai.Client, repository repository.HistoryRepositoryInterface, cache repository.CacheRepositoryInterface) *GeminiService {

	geminiServiceOnce.Do(func() {
		geminiRedisCacheKey = "gemini_" + config.EnvGeminiModel() + "_content"
	})

	return &GeminiService{client, repository, cache}
}

func (srv *GeminiService) ResetSession(ctx context.Context) (string, error) {
	if err := srv.cache.Delete(ctx, geminiRedisCacheKey); err != nil {
		return "", err
	}
	return "ok", nil
}

func (srv *GeminiService) PromptText(prompt *request.PromptText, ctx context.Context) (string, error) {
	resp, err := srv.Client.Models.GenerateContent(ctx, config.EnvGeminiModel(), genai.Text(prompt.Text), srv.getGenerateContentConfig())
	if err != nil {
		return "", err
	}

	return resp.Text(), nil
}

func (srv *GeminiService) Chat(prompt *request.PromptText, ctx context.Context) (string, error) {

	geminiContent, err := helper.GetTyped[[]*genai.Content](srv.cache, ctx, geminiRedisCacheKey)

	if errors.Is(err, redis.Nil) {
		geminiContent = make([]*genai.Content, 0)
	} else if err != nil {
		return "", err
	}

	if config.EnvRetrieveHistory() && len(geminiContent) == 0 {
		var histories = srv.GetAllByModelSource("gemini")

		for _, history := range histories {
			geminiContent = append(geminiContent, genai.NewContentFromText(history.UserQuestion, genai.RoleUser))
			geminiContent = append(geminiContent, genai.NewContentFromText(history.ModelAnswer, genai.RoleModel))
		}
	}

	chat, err := srv.Client.Chats.Create(ctx, config.EnvGeminiModel(), srv.getGenerateContentConfig(), geminiContent)

	if err != nil {
		return "", err
	}

	res, err := chat.SendMessage(ctx, genai.Part{Text: prompt.Text})

	if err != nil {
		return "", err
	}

	geminiContent = append(geminiContent, genai.NewContentFromText(prompt.Text, genai.RoleUser))
	geminiContent = append(geminiContent, genai.NewContentFromText(res.Text(), genai.RoleModel))

	if err := srv.cache.SetJSON(ctx, geminiRedisCacheKey, geminiContent); err != nil {
		return "", err
	}

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
