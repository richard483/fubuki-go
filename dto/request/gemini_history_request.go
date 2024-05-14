package request

type GeminiHistory struct {
	UserQuestion string `json:"question" binding:"required" validate:"required"`
	ModelAnswer  string `json:"answer" binding:"required" validate:"required"`
}

type UpdateGeminiHistory struct {
	ID           uint   `json:"id" binding:"required" validate:"required"`
	UserQuestion string `json:"question" binding:"required" validate:"required"`
	ModelAnswer  string `json:"answer" binding:"required" validate:"required"`
}
