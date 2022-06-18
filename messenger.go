package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

type SlackRequestBody struct {
	Text string `json:"text"`
}

func Message(status string) (string, error) {
	env := os.Getenv("ENV")
	webhookURL := viper.Get("SLACK_WEBHOOK_URL").(string)
	defaultStatus := viper.Get("DEFAULT_STATUS").(string)
	tag := ""
	if env != "prod" {
		tag = "[This is a test message.]"
	}
	var messsageContent string
	if status != defaultStatus {
		messsageContent = fmt.Sprintf("<!here> New update - case status is *%v* %v", status, tag)
	} else {
		messsageContent = fmt.Sprintf("Case status is *%v* %v", status, tag)
	}

	client := &http.Client{}
	slackBody, _ := json.Marshal(SlackRequestBody{Text: messsageContent})
	req, err := http.NewRequest("POST", webhookURL,
		bytes.NewBuffer(slackBody))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Fatalf("Request error %v\n", err)
		return "", err
	}

	log.Println("Sending Slack message ...")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Response error %v\n", err)
		return "", err
	}

	log.Println("Slack message sent.")
	defer resp.Body.Close()
	return "ok", nil
}
