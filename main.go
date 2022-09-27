package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tapiaw38/ecommerzeta-bot/handlers"
	"github.com/tapiaw38/ecommerzeta-bot/server"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error load env file")
	}

	PORT := os.Getenv("PORT")
	SLACK_TOKEN := os.Getenv("SLACK_TOKEN")
	SLACK_CHANNEL_FRONT := os.Getenv("SLACK_CHANNEL_FRONT")
	SLACK_CHANNEL_BACK := os.Getenv("SLACK_CHANNEL_BACK")
	SLACK_WEBHOOK_FRONT := os.Getenv("SLACK_WEBHOOK_FRONT")
	SLACK_WEBHOOK_BACK := os.Getenv("SLACK_WEBHOOK_BACK")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:              PORT,
		SlackToken:        SLACK_TOKEN,
		SlackChannelFront: SLACK_CHANNEL_FRONT,
		SlackChannelBack:  SLACK_CHANNEL_BACK,
		SlackWebhookFront: SLACK_WEBHOOK_FRONT,
		SlackWebhookBack:  SLACK_WEBHOOK_BACK,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Serve(handlers.BinderRoutes)
}
