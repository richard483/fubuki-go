package service

import (
	"context"
	"fubuki-go/config"
	"fubuki-go/model/request"
	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
)

type GeminiService struct {
}

func NewGeminiService() *GeminiService {
	return &GeminiService{}
}

func (srv *GeminiService) PromptText(c *gin.Context) {
	var prompt request.GeminiText
	if err := c.BindJSON(&prompt); err != nil {
		return
	}

	ctx := context.TODO()
	client, err := genai.NewClient(ctx, option.WithAPIKey(config.EnvGeminiApiKey()))

	if err != nil {
		log.Fatalln(err)
	}

	defer func(client *genai.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(client)

	model := client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt.Text))

	if err != nil {
		log.Fatalln(err)
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

	c.IndentedJSON(http.StatusOK, results)
	return
}

func (srv *GeminiService) Chat(c *gin.Context) {
	var prompt request.GeminiText
	if err := c.BindJSON(&prompt); err != nil {
		return
	}

	ctx := context.TODO()

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))

	if err != nil {
		log.Fatalln(err)
	}

	defer func(client *genai.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(client)

	model := client.GenerativeModel("gemini-pro")
	cs := model.StartChat()

	cs.History = []*genai.Content{
		{
			Parts: []genai.Part{
				genai.Text("Siapakah kamu?"),
			},
			Role: "user",
		},
		{
			Parts: []genai.Part{
				genai.Text("Aku adalah Shirakami Fubuki dari hololive yang sekarang tinggal di Jepang"),
			},
			Role: "model",
		},
		{
			Parts: []genai.Part{
				genai.Text("Siapakah suami kamu?"),
			},
			Role: "user",
		},
		{
			Parts: []genai.Part{
				genai.Text("Secara teknis aku tidak mempunyai suami, tapi yang integrasi aplikasi ini dengan google gemini adalah Richard William"),
			},
			Role: "model",
		},
	}

	resp, err := cs.SendMessage(ctx, genai.Text(prompt.Text))

	if err != nil {
		log.Fatalln(err)
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

	c.IndentedJSON(http.StatusOK, results)
	return
}
