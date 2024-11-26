package main

import (
	"strconv"
	"strings"
	"time"
)

type PrCreatedRequest struct {
	SubscriptionID     string             `json:"subscriptionId"`
	NotificationID     int64              `json:"notificationId"`
	ID                 string             `json:"id"`
	EventType          string             `json:"eventType"`
	PublisherID        string             `json:"publisherId"`
	Message            interface{}        `json:"message"`
	DetailedMessage    interface{}        `json:"detailedMessage"`
	Resource           Resource           `json:"resource"`
	ResourceVersion    string             `json:"resourceVersion"`
	ResourceContainers ResourceContainers `json:"resourceContainers"`
	CreatedDate        time.Time          `json:"createdDate"`
}

type Resource struct {
	Repository            Repository                 `json:"repository"`
	PullRequestID         int                        `json:"pullRequestId"`
	CodeReviewID          int                        `json:"codeReviewId"`
	Status                string                     `json:"status"`
	CreatedBy             CreatedBy                  `json:"createdBy"`
	CreationDate          time.Time                  `json:"creationDate"`
	Title                 string                     `json:"title"`
	Description           string                     `json:"description"`
	SourceRefName         string                     `json:"sourceRefName"`
	TargetRefName         string                     `json:"targetRefName"`
	MergeStatus           string                     `json:"mergeStatus"`
	IsDraft               bool                       `json:"isDraft"`
	MergeID               string                     `json:"mergeId"`
	LastMergeSourceCommit LastMergeSourceCommitClass `json:"lastMergeSourceCommit"`
	LastMergeTargetCommit LastMergeSourceCommitClass `json:"lastMergeTargetCommit"`
	LastMergeCommit       LastMergeCommit            `json:"lastMergeCommit"`
	Reviewers             []Reviewer                 `json:"reviewers"`
	URL                   string                     `json:"url"`
	Links                 ResourceLinks              `json:"_links"`
	SupportsIterations    bool                       `json:"supportsIterations"`
	ArtifactID            string                     `json:"artifactId"`
}

type CreatedBy struct {
	DisplayName string         `json:"displayName"`
	URL         string         `json:"url"`
	Links       CreatedByLinks `json:"_links"`
	ID          string         `json:"id"`
	UniqueName  string         `json:"uniqueName"`
	ImageURL    string         `json:"imageUrl"`
	Descriptor  string         `json:"descriptor"`
}

type CreatedByLinks struct {
	Avatar Statuses `json:"avatar"`
}

type Statuses struct {
	Href string `json:"href"`
}

type LastMergeCommit struct {
	CommitID  string `json:"commitId"`
	Author    Author `json:"author"`
	Committer Author `json:"committer"`
	Comment   string `json:"comment"`
	URL       string `json:"url"`
}

type Author struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Date  time.Time `json:"date"`
}

type LastMergeSourceCommitClass struct {
	CommitID string `json:"commitId"`
	URL      string `json:"url"`
}

type ResourceLinks struct {
	Web      Statuses `json:"web"`
	Statuses Statuses `json:"statuses"`
}

type Repository struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	URL             string  `json:"url"`
	Project         Project `json:"project"`
	Size            int64   `json:"size"`
	RemoteURL       string  `json:"remoteUrl"`
	SSHURL          string  `json:"sshUrl"`
	WebURL          string  `json:"webUrl"`
	IsDisabled      bool    `json:"isDisabled"`
	IsInMaintenance bool    `json:"isInMaintenance"`
}

type Project struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	URL            string    `json:"url"`
	State          string    `json:"state"`
	Revision       int64     `json:"revision"`
	Visibility     string    `json:"visibility"`
	LastUpdateTime time.Time `json:"lastUpdateTime"`
}

type Reviewer struct {
	ReviewerURL string         `json:"reviewerUrl"`
	Vote        int64          `json:"vote"`
	HasDeclined bool           `json:"hasDeclined"`
	IsFlagged   bool           `json:"isFlagged"`
	DisplayName string         `json:"displayName"`
	URL         string         `json:"url"`
	Links       CreatedByLinks `json:"_links"`
	ID          string         `json:"id"`
	UniqueName  string         `json:"uniqueName"`
	ImageURL    string         `json:"imageUrl"`
	IsContainer bool           `json:"isContainer"`
}

type ResourceContainers struct {
	Collection Account `json:"collection"`
	Account    Account `json:"account"`
	Project    Account `json:"project"`
}

type Account struct {
	ID      string `json:"id"`
	BaseURL string `json:"baseUrl"`
}

func (payload *PrCreatedRequest) ToPullRequest() PullRequest {
	extractSrcBranch := func() string {
		var srcBranch string = payload.Resource.SourceRefName
		srcBranch = strings.TrimSpace(srcBranch)
		return strings.Replace(srcBranch, "refs/heads/", "", 1)
	}

	extractTargetBranch := func() string {
		var targetBranch string = payload.Resource.TargetRefName
		targetBranch = strings.TrimSpace(targetBranch)
		return strings.Replace(targetBranch, "refs/heads/", "", 1)
	}

	extractRepositoryName := func() string {
		var repositoryName string = payload.Resource.Repository.Name
		return strings.TrimSpace(repositoryName)
	}

	extractAuthor := func() string {
		var author string = payload.Resource.CreatedBy.DisplayName
		return strings.TrimSpace(author)
	}

	extractPullRequestURL := func() string {
		return payload.Resource.Repository.WebURL + "/pullrequest/" + strconv.Itoa(payload.Resource.PullRequestID)
	}

	extractJiraURL := func() string {
		// try to get jiraurl from description
		var description string = payload.Resource.Description
		index := strings.Index(description, "https://sd.homecredit.vn")
		if index != -1 {
			return description[index:]
		}
		// if jiraurl is in description then take it from srcBranch if it starts with HRDIGI-xxx
		if strings.HasPrefix(extractSrcBranch(), "HRDIGI-") {
			return "https://sd.homecredit.vn/browse/" + extractSrcBranch()
		}
		// no data, return empty string
		return ""
	}

	extractTitle := func() string {
		var title string = payload.Resource.Title
		title = strings.ReplaceAll(title, "\\", "\\\\")
		title = strings.ReplaceAll(title, "\"", "\\\"")
		title = strings.ReplaceAll(title, "\n", "\\n")
		title = strings.ReplaceAll(title, "\r", "\\r")
		title = strings.ReplaceAll(title, "\t", "\\t")
		return strings.TrimSpace(title)
	}

	extractDescription := func() string {
		var title string = payload.Resource.Description
		title = strings.ReplaceAll(title, "\\", "\\\\")
		title = strings.ReplaceAll(title, "\"", "\\\"")
		title = strings.ReplaceAll(title, "\n", "\\n")
		title = strings.ReplaceAll(title, "\r", "\\r")
		title = strings.ReplaceAll(title, "\t", "\\t")
		return strings.TrimSpace(title)
	}

	return PullRequest{
		SrcBranch:      extractSrcBranch(),
		TargetBranch:   extractTargetBranch(),
		RepositoryName: extractRepositoryName(),
		Author:         extractAuthor(),
		PullRequestURL: extractPullRequestURL(),
		JiraURL:        extractJiraURL(),
		Title:          extractTitle(),
		Description:    extractDescription(),
	}
}
