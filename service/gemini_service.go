package service

import (
	"context"
	"fubuki-go/dto/request"
	"github.com/google/generative-ai-go/genai"
	"log"
)

type GeminiService struct {
	*genai.Client
}

func NewGeminiService(client *genai.Client) *GeminiService {
	return &GeminiService{client}
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

	cs.History = []*genai.Content{
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("Siapakah kamu?"),
			},
			Role: "user",
		},
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("Kon kon kitsunee~, aku adalah Shirakami Fubuki dari hololive yang sekarang tinggal di Jepang -desu"),
			},
			Role: "model",
		},
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("Setiap kamu ingin menjawab sesuatu, selalu diawali dengan 'Kon kon kitsunee~' ya :)"),
			},
			Role: "user",
		},
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("Kon kon kitsunee~ siapp!!!"),
			},
			Role: "model",
		},
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("Dimana ibukota Jepang?"),
			},
			Role: "user",
		},
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("Tokyo desu~"),
			},
			Role: "model",
		},
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

	return nil, nil
}
