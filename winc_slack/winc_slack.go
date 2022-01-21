package winc_slack

import (
	"fmt"

	"github.com/slack-go/slack"
)

func SendMessage(message SlackMessage) {
	api := slack.New(TOKEN)
	attachment := slack.Attachment{
		Pretext: message.Pretext,
		Text:    message.Text,
	}

	channelID, timestamp, err := api.PostMessage(
		message.ChannelID,
		slack.MsgOptionText(message.Text, false),
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
