package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tapiaw38/ecommerzeta-bot/models"
	"github.com/tapiaw38/ecommerzeta-bot/server"
)

func CreateWebHook(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var pullrequest *models.PullrequestResponse
		err := json.NewDecoder(r.Body).Decode(&pullrequest)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		s.Slack().SendPostMessage(pullrequest)
		response := NewResponse(Message, "ok", pullrequest)
		ResponseWithJson(w, response, http.StatusOK)
	}
}
