package server

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/tapiaw38/ecommerzeta-bot/utils"
)

type Server interface {
	Config() *Config
	Slack() *utils.SlackConfig
}

type Config struct {
	Port              string
	SlackToken        string
	SlackChannelFront string
	SlackChannelBack  string
	SlackWebhookFront string
	SlackWebhookBack  string
}

type Broker struct {
	config *Config
	router *mux.Router
	slack  *utils.SlackConfig
}

func (b *Broker) Config() *Config {
	return b.config
}

func (b *Broker) Slack() *utils.SlackConfig {
	return b.slack
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
		slack: utils.NewSlack(
			config.SlackToken,
			config.SlackChannelFront,
			config.SlackChannelBack,
			config.SlackWebhookFront,
			config.SlackWebhookBack,
		),
	}

	return broker, nil
}

func (b *Broker) Serve(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	handler := cors.AllowAll().Handler(b.router)

	log.Println("listening on port", b.config.Port)

	if err := http.ListenAndServe(":"+b.config.Port, handler); err != nil {
		log.Fatal(err)
	}
}
