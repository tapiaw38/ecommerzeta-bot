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
		log.Fatal("Error load env file")
	}

	PORT := os.Getenv("PORT")
	SLACK_TOKEN := os.Getenv("SLACK_TOKEN")
	SLACK_CHANNEL := os.Getenv("SLACK_CHANNEL")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:         PORT,
		SlackToken:   SLACK_TOKEN,
		SlackChannel: SLACK_CHANNEL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Serve(handlers.BinderRoutes)
}
