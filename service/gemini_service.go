package service

import (
	"context"
	"fubuki-go/dto/request"
	repository "fubuki-go/repository_interface"
	"github.com/google/generative-ai-go/genai"
	"log"
)

type GeminiService struct {
	*genai.Client
	repository.GeminiHistoryRepositoryInterface
}

func NewGeminiService(client *genai.Client, repository repository.GeminiHistoryRepositoryInterface) *GeminiService {
	return &GeminiService{client, repository}
}

func (srv *GeminiService) PromptText(prompt *request.GeminiText) (error, *[]string) {
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

func (srv *GeminiService) Chat(prompt *request.GeminiText) (error, *[]string) {
	ctx := context.TODO()

	model := srv.Client.GenerativeModel("gemini-pro")
	cs := model.StartChat()

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

	resp, err := cs.SendMessage(ctx, genai.Text(prompt.Text))

	if err != nil {
		return err, nil
	}
	log.Println(resp)

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
