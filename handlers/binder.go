package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tapiaw38/ecommerzeta-bot/server"
)

func BinderRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/webhook", CreateWebHook(s)).Methods(http.MethodPost)
}
