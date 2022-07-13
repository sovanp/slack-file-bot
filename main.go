package main

import (
	"os"

	"github.com/slack-go/slack"
)

func main() {

	// os.Setenv("SLACK_BOT_TOKEN", "xoxb-3761425391125-3796134192258-kJQNBU4uRv8EnsbbNhxEBvFz")
	// os.Setenv("CHANNEL_ID", "")

	//Variables to store token, channel id, and files
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{""}

}
