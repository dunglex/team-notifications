package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AdaptiveCardRequest struct {
	Type        string                   `json:"type"`
	Attachments []AdaptiveCardAttachment `json:"attachments"`
}

type AdaptiveCardAttachment struct {
	ContentType string         `json:"contentType"`
	ContentUrl  *string        `json:"contentUrl"`
	Content     []AdaptiveCard `json:"content"`
}

type AdaptiveCard struct {
	Schema  string             `json:"$schema"`
	Type    string             `json:"type"`
	Version string             `json:"version"`
	Body    []AdaptiveCardBody `json:"body"`
}

type AdaptiveCardBody struct {
	Type           string `json:"type"`
	Title          string `json:"title"`
	Text           string `json:"text"`
	Url            string `json:"url"`
	JiraUrl        string `json:"jiraUrl"`
	SourceBranch   string `json:"srcBranch"`
	TargetBranch   string `json:"targetBranch"`
	RepositoryName string `json:"repository"`
	Author         string `json:"author"`
}

func (card *AdaptiveCard) sendAdaptiveCard(webhookUrl string, dumpJson bool) error {
	var attachment = AdaptiveCardAttachment{
		ContentType: "application/vnd.microsoft.card.adaptive",
		Content:     []AdaptiveCard{*card},
	}

	var request = AdaptiveCardRequest{
		Type:        "message",
		Attachments: []AdaptiveCardAttachment{attachment},
	}

	payloadBytes, err := json.Marshal(request)
	if err != nil {
		return err
	}

	if dumpJson {
		fmt.Println("Sending message: " + string(payloadBytes))
	}

	req, err := http.NewRequest("POST", webhookUrl, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	return fmt.Errorf("failed to send Adaptive Card, status code: %d", resp.StatusCode)
}
