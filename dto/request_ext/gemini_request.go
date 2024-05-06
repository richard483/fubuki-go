package request

type GeminiPart struct {
	Text string `json:"text,omitempty" binding:"required"`
}

type GeminiContent struct {
	Role  string        `json:"role,omitempty" binding:"required"`
	Parts *[]GeminiPart `json:"parts,omitempty" binding:"required"`
}

type GeminiContents struct {
	Contents *[]GeminiContent `json:"contents,omitempty" binding:"required"`
}
