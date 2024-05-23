package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func NewResponse(statusCode int, message string, data interface{}) *Response {
	return &Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}

func (r *Response) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	json.NewEncoder(w).Encode(r)
}
