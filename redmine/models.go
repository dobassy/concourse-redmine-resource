package redmine

import (
	"net/http"
	"net/url"
)

type Client struct {
	apikey     string
	URL        *url.URL
	HTTPClient *http.Client
}

type PostIssueRequest struct {
	Issue PostIssueContent `json:"issue"`
}

type PostIssueContent struct {
	ProjectId   int    `json:"project_id,omitempty"`
	TrackerId   int    `json:"tracker_id,omitempty"`
	StatusId    int    `json:"status_id,omitempty"`
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
	Id             int           `json:"id"`
	Project        IdNamePair    `json:"project"`
	Tracker        IdNamePair    `json:"tracker"`
	Status         IdNamePair    `json:"status"`
	Priority       IdNamePair    `json:"priority"`
	Author         IdNamePair    `json:"author"`
	AssignedTo     IdNamePair    `json:"assigned_to,omitempty"`
	Category       IdNamePair    `json:"category,omitempty"`
	FixedVersion   IdNamePair    `json:"fixed_version,omitempty"`
	Parent         Id            `json:"parent,omitempty"`
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

type Id struct {
	Id int `json:"id"`
}

type IdNamePair struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CustomField struct {
	Id    int         `json:"id"`
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}
