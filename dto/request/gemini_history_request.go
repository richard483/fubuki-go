package request

type History struct {
	UserQuestion string `json:"question" binding:"required" validate:"required"`
	ModelAnswer  string `json:"answer" binding:"required" validate:"required"`
	ModelSource  string `json:"model_source" binding:"required" validate:"required"`
}

type UpdateHistory struct {
	ID           uint   `json:"id" binding:"required" validate:"required"`
	UserQuestion string `json:"question" validate:"required"`
	ModelAnswer  string `json:"answer" validate:"required"`
	ModelSource  string `json:"model_source" validate:"required"`
}
