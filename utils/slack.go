package utils

import (
	"log"
	"strings"

	"github.com/ashwanthkumar/slack-go-webhook"
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
		"Ignacio Pieve Roiger",
		"ildemaro.carrasco",
		"William Buezas Madsen",
	}
)

type SlackConfig struct {
	slackToken        string
	slackChannelFront string
	slackChannelBack  string
	slackWebhookFront string
	slackWebhookBack  string
}

func NewSlack(
	slackToken string,
	slackChannelFront string,
	slackChannelBack string,
	slackWebhookFront string,
	slackWebhookBack string,
) *SlackConfig {
	return &SlackConfig{
		slackToken:        slackToken,
		slackChannelFront: slackChannelFront,
		slackChannelBack:  slackChannelBack,
		slackWebhookFront: slackWebhookFront,
		slackWebhookBack:  slackWebhookBack,
	}
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
		log.Printf("The user is not in the ecommerce list, name: %s, id: %s, repo: %s", pullrequest.Actor.DisplayName, pullrequest.Actor.AccountId, pullrequest.Repository.Name)
		return
	}

	var channel string
	var webhook string

	switch pullrequest.Repository.Name {
	case "bigbox":
		channel = s.slackChannelBack
		webhook = s.slackWebhookBack
	case "Bigbox-Frontend":
		channel = s.slackChannelFront
		webhook = s.slackWebhookFront
	default:
		channel = "#la-ecommerzeta"
	}

	attachment := slack.Attachment{}
	attachment.AuthorIcon = &pullrequest.Actor.Links.Avatar.Href
	attachment.AuthorName = &pullrequest.Actor.DisplayName
	attachment.Title = &pullrequest.Pullrequest.Title
	attachment.TitleLink = &pullrequest.Pullrequest.Links.Html.Href
	attachment.Text = &pullrequest.Pullrequest.Description
	attachment.AddField(slack.Field{Title: "Reviewers"})

	if len(pullrequest.Pullrequest.Reviewers) > 0 {
		for _, r := range pullrequest.Pullrequest.Reviewers {
			reviewer := "@" + r.DisplayName
			attachment.AddField(slack.Field{Value: reviewer})
		}
	} else {
		attachment.AddField(slack.Field{Value: "No reviewers"})
	}

	payload := slack.Payload{
		Text:        "🎉​ `" + pullrequest.Actor.DisplayName + "` has created a new pull request",
		Username:    "Ecommerzeta",
		Channel:     channel,
		IconEmoji:   ":gorro_de_fiesta:",
		Attachments: []slack.Attachment{attachment},
	}

	err := slack.Send(webhook, "", payload)
	if len(err) > 0 {
		log.Printf("error: %s\n", err)
	}
	log.Printf("Pullrequest, name %s, id %s", pullrequest.Actor.DisplayName, pullrequest.Actor.AccountId)
}
