package utils

import (
	"log"

	"github.com/slack-go/slack"
)

type SlackConfig struct {
	slackToken   string
	slackChannel string
}

func NewSlack(slackToken string, slackChannel string) *SlackConfig {
	return &SlackConfig{
		slackToken:   slackToken,
		slackChannel: slackChannel,
	}
}

func (s SlackConfig) SlackClient() *slack.Client {
	return slack.New(s.slackToken)
}

func (s SlackConfig) SendPostMessage(message string) {

	_, _, err := s.SlackClient().PostMessage(
		s.slackChannel,
		slack.MsgOptionText(
			message,
			false),
	)

	if err != nil {
		log.Printf("%s\n", err)
	}

	log.Println("Message successfully")
}
