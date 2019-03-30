package redmine

import (
	"net/http"
	"net/url"
)

type Client interface {
	GetIssues() (*IssuesResponse, error)
	CreateIssue(issue PostIssueContent) (*PostIssueResponse, error)
}

type client struct {
	apikey     string
	URL        *url.URL
	HTTPClient *http.Client
}

type PostIssueRequest struct {
	Issue PostIssueContent `json:"issue"`
}

type PostIssueContent struct {
	ProjectID   int    `json:"project_id,omitempty"`
	TrackerID   int    `json:"tracker_id,omitempty"`
	StatusID    int    `json:"status_id,omitempty"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

type PostIssueResponse struct {
	Issue      Issue `json:"issue"`
	TotalCount uint  `json:"total_count"`
	Offset     uint  `json:"offset"`
	Limit      uint  `json:"limit"`
}

type IssuesResponse struct {
	Issues     []Issue `json:"issues"`
	TotalCount uint    `json:"total_count"`
	Offset     uint    `json:"offset"`
	Limit      uint    `json:"limit"`
}

type Issue struct {
	ID             int           `json:"id"`
	Project        IDNamePair    `json:"project"`
	Tracker        IDNamePair    `json:"tracker"`
	Status         IDNamePair    `json:"status"`
	Priority       IDNamePair    `json:"priority"`
	Author         IDNamePair    `json:"author"`
	AssignedTo     IDNamePair    `json:"assigned_to,omitempty"`
	Category       IDNamePair    `json:"category,omitempty"`
	FixedVersion   IDNamePair    `json:"fixed_version,omitempty"`
	Parent         ID            `json:"parent,omitempty"`
	Subject        string        `json:"subject"`
	Description    string        `json:"description"`
	StartDate      string        `json:"start_date"`
	DueDate        string        `json:"due_date"`
	DoneRatio      int           `json:"done_ratio,omitempty"`
	IsPrivate      bool          `json:"is_private,omitempty"`
	EstimatedHours int           `json:"estimated_hours,omitempty"`
	CustomFields   []CustomField `json:"custom_fields,omitempty"`
	CreatedOn      string        `json:"created_on"`
	UpdatedOn      string        `json:"updated_on"`
	ClosedOn       string        `json:"closed_on"`
}

type ID struct {
	ID int `json:"id"`
}

type IDNamePair struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CustomField struct {
	ID    int         `json:"id"`
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}
