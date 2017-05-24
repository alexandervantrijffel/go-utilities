package slackposter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type SlackPayload struct {
	Text      string `json:"text"`
	IconEmoji string `json:"icon_emoji"`
}

func PostToSlack(slackWebhookURL string, iconEmoji string, text string) error {
	return SlackPayload{IconEmoji: iconEmoji, Text: text}.Post(slackWebhookURL)
}

func SPostToSlack(slackWebhookURL string, iconEmoji string, formatString string, formatParameters ...interface{}) error {
	return PostToSlack(slackWebhookURL, iconEmoji, fmt.Sprintf(formatString, formatParameters))
}

func (payload SlackPayload) Post(slackWebhookURL string) error {
	client := &http.Client{}
	marshalled, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	response, err := client.Post(slackWebhookURL, "application/json", bytes.NewReader(marshalled))
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return fmt.Errorf(fmt.Sprintf("Did not receive expected HTTP 200 from Slack. Received: %s (%d)",
			response.Status, response.StatusCode))
	}
	return nil
}
