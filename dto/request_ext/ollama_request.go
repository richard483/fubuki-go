package request_ext

type OllamaGenerateRequest struct {
	Model  string `json:"model,omitempty" binding:"required"`
	Prompt string `json:"prompt,omitempty" binding:"required"`
	Stream string `json:"Stream,omitempty" binding:"required"`
}
