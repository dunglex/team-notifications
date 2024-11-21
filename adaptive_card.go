package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AdaptiveCardRequest struct {
	Type        string         `json:"type"`
	Attachments []AdaptiveCard `json:"attachments"`
}

type AdaptiveCard struct {
	Type    string             `json:"type"`
	Version string             `json:"version"`
	Body    []AdaptiveCardBody `json:"body"`
}

type AdaptiveCardBody struct {
	Type     string `json:"type"`
	Text     string `json:"text"`
	Wrap     bool   `json:"wrap,omitempty"`
	Size     string `json:"size,omitempty"`
	Weight   string `json:"weight,omitempty"`
	Spacing  string `json:"spacing,omitempty"`
	IsSubtle bool   `json:"isSubtle,omitempty"`
}

func (card *AdaptiveCard) sendAdaptiveCard(webhookUrl string, dumpJson bool) error {
	var request = AdaptiveCardRequest{
		Type:        "message",
		Attachments: []AdaptiveCard{*card},
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
