package request_ext

type OllamaGenerateRequest struct {
	Model  string `json:"model" binding:"required"`
	Prompt string `json:"prompt" binding:"required"`
	Stream bool   `json:"stream" binding:"required" default:"false"`
}

type OllamaChatRequest struct {
	Model    string          `json:"model" binding:"required"`
	Messages []OllamaMessage `json:"messages" binding:"required"`
	Stream   bool            `json:"stream" binding:"required" default:"false"`
}

type OllamaMessage struct {
	Role    string `json:"role" binding:"required"`
	Content string `json:"content" binding:"required"`
}
