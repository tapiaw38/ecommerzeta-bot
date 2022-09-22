package utils

import (
	"log"
	"strings"

	"github.com/slack-go/slack"
	"github.com/tapiaw38/ecommerzeta-bot/models"
)

var (
	SEND_NEDDED = []string{
		"Walter Tapia",
		"Federico Poncela",
		"Julian Liendo",
		"Daniel Leonardo Villajuan",
		"Julián Gregori Battista",
		"Elias Velardez",
		"Anibal Reinoso",
		"Ignacio Pieve Roiger",
		"ildemaro.carrasco",
	}
)

type SlackConfig struct {
	slackToken        string
	slackChannelFront string
	slackChannelBack  string
}

func NewSlack(slackToken string, slackChannelFront string, slackChannelBack string) *SlackConfig {
	return &SlackConfig{
		slackToken:        slackToken,
		slackChannelFront: slackChannelFront,
		slackChannelBack:  slackChannelBack,
	}
}

func (s SlackConfig) SlackClient() *slack.Client {
	return slack.New(s.slackToken)
}

func shouldCheckUser(displayName string) bool {
	for _, u := range SEND_NEDDED {
		if strings.Contains(displayName, u) {
			return true
		}
	}
	return false
}

func (s SlackConfig) SendPostMessage(pullrequest *models.PullrequestResponse) {

	if !shouldCheckUser(pullrequest.Actor.DisplayName) {
		log.Printf("The user is not in the ecommerce list, name: %s, id: %s", pullrequest.Actor.NickName, pullrequest.Actor.AccountId)
		return
	}

	var channel string

	switch pullrequest.Repository.Name {
	case "bigbox":
		channel = s.slackChannelBack
	case "Bigbox_Frontend":
		channel = s.slackChannelFront
	default:
		channel = "general"
	}

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
		channel,
		slack.MsgOptionAttachments(attachment),
	)

	if err != nil {
		log.Printf("%s\n", err)
	}

	log.Printf("Pullrequest, name %s, id %s", pullrequest.Actor.NickName, pullrequest.Actor.AccountId)
}
