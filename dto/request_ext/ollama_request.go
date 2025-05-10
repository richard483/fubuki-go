package request_ext

type OllamaGenerateRequest struct {
	Model  string `json:"model" binding:"required"`
	Prompt string `json:"prompt" binding:"required"`
	Stream bool   `json:"stream" binding:"required"`
}
