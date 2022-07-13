package main

import (
	"os"

	"github.com/slack-go/slack"
)

func main() {
	//Variables to store token, channel id, and files
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{""}
}
