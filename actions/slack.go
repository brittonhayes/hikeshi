package actions

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/brittonhayes/hikeshi/models"
	"github.com/gobuffalo/buffalo/worker"
	"github.com/slack-go/slack"
)

//var w worker.Worker

// TODO Implement background worker to send messages

//func init() {
//	w = App().Worker // Get a ref to the previously defined Worker
//	err := w.Register("send_slack_message", func(args worker.Args) error {
//
//		// Convert args into strings
//		webhookURL := fmt.Sprintf("%s", args["webhook_url"])
//		message := fmt.Sprintf("%s", args["message"])
//
//		// Send the request
//		err := SendSlackNotification(baseURL, webhookURL, message)
//		if err != nil {
//			return err
//		}
//
//		return err
//	})
//	if err != nil {
//		log.Print(err)
//	}
//}

type Slack struct {
	Request SlackRequest
	Job     worker.Job
}

type SlackRequest struct {
	Text string `json:"text"`
}

type SlackPayload map[string]interface{}

// SendSlackNotification will post to an 'Incoming Webook' url setup in Slack Apps. It accepts
// some text and the slack channel is saved within Slack.
func SendSlackNotification(baseURL string, webhookURL string, incident models.Incident) error {
	b := NewBlockMessage(baseURL, incident)
	err := slack.PostWebhook(webhookURL, b)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func NewBlockMessage(baseURL string, i models.Incident) *slack.WebhookMessage {

	headerText := slack.NewTextBlockObject("plain_text", ":bell: New Incident Submitted", true, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	titleText := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Title:* `%s`", i.Title), false, false)
	titleSection := slack.NewSectionBlock(titleText, nil, nil)

	dateText := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Date:* `%s`", i.Date.String()), false, false)
	dateSection := slack.NewSectionBlock(dateText, nil, nil)

	severityText := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Severity:* `%s`", i.Severity), false, false)
	severitySection := slack.NewSectionBlock(severityText, nil, nil)

	dividerSection := slack.NewDividerBlock()

	accessoryText := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("For security reasons, further details can be <%s/incidents/%s|found in Hikeshi>", baseURL, i.ID), false, false)
	accessorySection := slack.NewSectionBlock(accessoryText, nil, nil)

	//accessorySection := slack.NewSectionBlock(accessoryText, nil, accessoryBody)
	blocks := &slack.Blocks{BlockSet: []slack.Block{headerSection, titleSection, dateSection, severitySection, dividerSection, accessorySection}}
	msg := &slack.WebhookMessage{
		IconEmoji: "fire",
		Blocks:    blocks,
	}

	b, _ := json.Marshal(msg)
	fmt.Print(string(b))

	return msg
}
