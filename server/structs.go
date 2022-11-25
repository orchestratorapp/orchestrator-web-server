package server

import "net/http"

// The Router handler
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

// The error response struct
type ErrorResponse struct {
	Status  int    `json:"status_code"`
	Message string `json:"message"`
}

func BuildErrorResponse(status int, err error) *ErrorResponse {
	return &ErrorResponse{
		Status:  status,
		Message: err.Error(),
	}
}
