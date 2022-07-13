package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	// load .env file from path
	err_env := godotenv.Load(".env")
	if err_env != nil {
		log.Fatalf("Error loading .env file")
	}

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))						// Store Bot Token
	channelArr := []string{os.Getenv("CHANNEL_ID")}						// Store Channel ID
	fileArr := []string{"dummy.pdf"}									// Store Filename

	//Loop over fileArr
	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters {
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}
