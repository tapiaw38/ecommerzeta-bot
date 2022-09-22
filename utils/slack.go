package utils

import (
	"log"

	"github.com/slack-go/slack"
	"github.com/tapiaw38/ecommerzeta-bot/models"
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

func (s SlackConfig) SendPostMessage(pullrequest *models.PullrequestResponse) {

	attachment := slack.Attachment{
		Color:      "#36a64f",
		Pretext:    "🎉 `" + pullrequest.Actor.DisplayName + "` has created a new pull request",
		AuthorName: pullrequest.Actor.DisplayName,
		AuthorIcon: pullrequest.Actor.Links.Avatar.Href,
		Title:      pullrequest.Pullrequest.Title,
		TitleLink:  pullrequest.Pullrequest.Links.Html.Href,
		Text:       pullrequest.Pullrequest.Description,
	}

	_, _, err := s.SlackClient().PostMessage(
		s.slackChannel,
		slack.MsgOptionAttachments(attachment),
	)

	if err != nil {
		log.Printf("%s\n", err)
	}

	log.Println("Message successfully")
}
