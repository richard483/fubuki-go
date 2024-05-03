package request

type GeminiPart struct {
	Text string `json:"text,omitempty" binding:"required"`
}

type GeminiContent struct {
	Parts []GeminiPart `json:"parts,omitempty" binding:"required"`
	Role  string       `json:"role,omitempty" binding:"required"`
}

type GeminiCandidate struct {
	Content GeminiContent  `json:"content,omitempty" binding:"required"`
	X       map[string]any `json:"-"`
}

type GeminiCandidates struct {
	Candidates []GeminiCandidate `json:"candidates,omitempty" binding:"required"`
}
