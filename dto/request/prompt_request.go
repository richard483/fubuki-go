package request

type PromptText struct {
	Text  string `json:"text,omitempty" binding:"required" validate:"required"`
	Model string `json:"model,omitempty" binding:"required"`
}
