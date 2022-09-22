package handlers

import (
	"encoding/json"
	"net/http"
)

const (
	Error   = "error"
	Message = "message"
)

type Response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Response    interface{} `json:"response"`
}

func NewResponse(messageType string, message string, response interface{}) Response {
	return Response{
		MessageType: messageType,
		Message:     message,
		Response:    response,
	}
}

func ResponseWithJson(w http.ResponseWriter, response Response, statusCode int) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
