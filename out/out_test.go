package out

import (
	"fmt"
	"testing"

	"github.com/dobassy/concourse-redmine-resource"
	"github.com/dobassy/concourse-redmine-resource/redmine"
)

var basicRequest = Request{
	Source: resource.Source{
		URI:      "http://localhost/",
		Apikey:   "dummykey",
		Insecure: false,
	},
	Params: Params{
		ProjectID: 1,
		TrackerID: 1,
		StatusID:  1,
		Subject:   "stdin subject",
	},
}

type redmineMock struct {
}

func NewRedmineMock() redmine.Client {
	return &redmineMock{}
}

func (m *redmineMock) GetIssues() (*redmine.IssuesResponse, error) {
	return &redmine.IssuesResponse{}, nil
}
func (m *redmineMock) CreateIssue(issue redmine.PostIssueContent) (*redmine.PostIssueResponse, error) {
	return &redmine.PostIssueResponse{}, nil
}

func TestPut(t *testing.T) {
	client := NewRedmineMock()
	res, err := Put(client, basicRequest)
	if err != nil {
		t.Fatalf("failed to Put(): %+v", err)
	}

	if !res.Version.Timestamp.IsZero() {
		t.Fatalf("Timestamp may not be correct: %+v", res)
	}
	fmt.Printf("%+v", res)
}
