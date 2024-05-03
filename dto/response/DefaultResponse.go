package response

type DefaultResponse struct {
	StatusCode int    `json:"status,omitempty"`
	Message    string `json:"message,omitempty"`
	Data       any    `json:"data,omitempty"`
	Error      string `json:"error,omitempty"`
}
