package request

type GeminiText struct {
	Text string `json:"text,omitempty" binding:"required"`
}
