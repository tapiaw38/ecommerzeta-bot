package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tapiaw38/ecommerzeta-bot/models"
	"github.com/tapiaw38/ecommerzeta-bot/server"
)

func CreateWebHook(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var data interface{}

		err := json.NewDecoder(r.Body).Decode(&data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		log.Println(data)

		response := NewResponse(Message, "ok", "ok")
		ResponseWithJson(w, response, http.StatusOK)
	}
}

func CreatePullRequest(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var pullrequest *models.PullRequest
		err := json.NewDecoder(r.Body).Decode(&pullrequest)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		message := pullrequest.PullRequestFormat()
		s.Slack().SendPostMessage(message)
		response := NewResponse(Message, "ok", pullrequest)
		ResponseWithJson(w, response, http.StatusOK)
	}
}
