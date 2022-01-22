package winc_slack

import (
	"fmt"

	"github.com/slack-go/slack"
)

func SendMessage(text string, pretext string, channelId string) {
	api := slack.New(TOKEN)
	attachment := slack.Attachment{
		Pretext: pretext,
		Text:    text,
	}

	channelID, timestamp, err := api.PostMessage(
		channelId,
		slack.MsgOptionText(text, false),
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
