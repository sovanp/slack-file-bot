# Slack File Bot

A simple Slack Bot that uploads any file, given a specified channel

## How to Use

1. Add the Slack Bot to your channel

2. Run the commands `go build` and then `source .env && go run main.go` to upoad the file

## Setup

1. Clone the repo `git clone <repo>`

2. Install dependencies `go get "github.com/joho/godotenv"` and `go get "github.com/slack-go/slack"`

3. Visit https://api.slack.com/apps to configure Slack Bot

   - Enable Socket Mode

   - Install app to Slack Workspace to obtain `SLACK_BOT_TOKEN`

   - Right-click on your channel and select `View channel details` to obtain `CHANNEL_ID`

Your `.env` file should be formatted as follows:

```
SLACK_BOT_TOKEN=
CHANNEL_ID=
```

## Setup

- Extend functionality to upload multiple files

- Extend functionality to upload in different channels
