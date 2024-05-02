package request

type CreateGeminiHistory struct {
	UserQuestion string `json:"question,omitempty" binding:"required" validate:"required"`
	ModelAnswer  string `json:"answer,omitempty" binding:"required" validate:"required"`
}

type UpdateGeminiHistory struct {
	UserQuestion string `json:"question,omitempty" binding:"required" validate:"required"`
	ModelAnswer  string `json:"answer,omitempty" binding:"required" validate:"required"`
}
