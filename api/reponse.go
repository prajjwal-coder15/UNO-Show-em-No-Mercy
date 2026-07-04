package api

// Response is the common response format.
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// ErrorResponse is returned on failures.
type ErrorResponse struct {
	Error string `json:"error"`
}