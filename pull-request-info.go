package main

import "fmt"

type PullRequest struct {
	SrcBranch      string `json:"sourceRefName"`
	TargetBranch   string `json:"targetRefName"`
	RepositoryName string `json:"repository"`
	Author         string `json:"createdBy"`
	PullRequestURL string `json:"url"`
	JiraURL        string `json:"jiraUrl"`
	Title          string `json:"title"`
	Description    string `json:"description"`
}

func (pr *PullRequest) CreateAdaptiveCard() AdaptiveCard {
	return AdaptiveCard{
		Type:    "AdaptiveCard",
		Version: "1.2",
		Body: []AdaptiveCardBody{
			{
				Type:   "TextBlock",
				Text:   fmt.Sprintf("New Pull Request Created by %s", pr.Author),
				Size:   "Large",
				Weight: "Bolder",
			},
			{
				Type: "TextBlock",
				Text: fmt.Sprintf("Title: %s", pr.Title),
				Wrap: true,
			},
			{
				Type: "TextBlock",
				Text: fmt.Sprintf("Description: %s", pr.Description),
				Wrap: true,
			},
			{
				Type: "TextBlock",
				Text: fmt.Sprintf("Source Branch: %s", pr.SrcBranch),
				Wrap: true,
			},
			{
				Type: "TextBlock",
				Text: fmt.Sprintf("Target Branch: %s", pr.TargetBranch),
				Wrap: true,
			},
			{
				Type: "TextBlock",
				Text: fmt.Sprintf("Repository: %s", pr.RepositoryName),
				Wrap: true,
			},
			{
				Type: "TextBlock",
				Text: fmt.Sprintf("Pull Request URL: %s", pr.PullRequestURL),
				Wrap: true,
			},
			{
				Type: "TextBlock",
				Text: fmt.Sprintf("JIRA URL: %s", pr.JiraURL),
				Wrap: true,
			},
		},
	}
}
