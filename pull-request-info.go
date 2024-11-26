package main

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
		Schema:  "http://adaptivecards.io/schemas/adaptive-card.json",
		Type:    "AdaptiveCard",
		Version: "1.2",
		Body: []AdaptiveCardBody{
			{
				Type:           "TextBlock",
				Title:          pr.Title,
				Text:           pr.Author,
				Url:            pr.PullRequestURL,
				JiraUrl:        pr.JiraURL,
				SourceBranch:   pr.SrcBranch,
				TargetBranch:   pr.TargetBranch,
				RepositoryName: pr.RepositoryName,
				Author:         pr.Author,
			},
		},
	}
}
