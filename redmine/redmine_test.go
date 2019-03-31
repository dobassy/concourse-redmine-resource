package redmine

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// This handler is used to test the func GetIssues()
var successResponseHandler = func(w http.ResponseWriter, r *http.Request) {
	json := `
		{
			"issues": [
				{
					"id": 10,
					"project": {
						"id": 1,
						"name": "Project 1"
					},
					"tracker": {
						"id": 1,
						"name": "Bug"
					},
					"status": {
						"id": 2,
						"name": "New"
					},
					"priority": {
						"id": 3,
						"name": "Normal"
					},
					"author": {
						"id": 9,
						"name": "Admin"
					},
					"subject": "Ticket title",
					"description": "description content",
					"start_date": "2019-03-24",
					"due_date": null,
					"done_ratio": 0,
					"is_private": false,
					"estimated_hours": null,
					"created_on": "2019-03-24T16:38:38Z",
					"updated_on": "2019-03-24T16:38:38Z",
					"closed_on": null
				}
			],
			"total_count": 1,
			"offset": 0,
			"limit": 25
		}
		`
	fmt.Fprintf(w, json)
}

// This handler is used to test the func CreateIssue()
var successCreateIssueResponseHandler = func(w http.ResponseWriter, r *http.Request) {
	json := `
		{
			"issue": {
				"id": 11,
				"project": {
					"id": 1,
					"name": "Project 1"
				},
				"tracker": {
					"id": 1,
					"name": "Support"
				},
				"status": {
					"id": 2,
					"name": "New"
				},
				"priority": {
					"id": 3,
					"name": "Normal"
				},
				"author": {
					"id": 9,
					"name": "Admin"
				},
				"subject": "Ticket title 2",
				"description": "description content",
				"start_date": "2019-03-24",
				"due_date": null,
				"done_ratio": 0,
				"is_private": false,
				"estimated_hours": null,
				"created_on": "2019-03-24T16:38:38Z",
				"updated_on": "2019-03-24T16:38:38Z",
				"closed_on": null
			}
		}
		`
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, json)
}

var badStatusResponseHandler = func(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Found.", http.StatusNotFound)
}

func TestNewClient(t *testing.T) {
	t.Run("when redmineurl key is blank", func(t *testing.T) {
		client, _ := NewClient(
			"",
			"dummyapikey",
			false,
		)

		if client != nil {
			t.Fatalf("failed validation: Initialization has succeeded even though there are no arguments: %v", client)
		}
	})

	t.Run("when apikey key is blank", func(t *testing.T) {
		client, _ := NewClient(
			"http://localhost",
			"",
			false,
		)

		if client != nil {
			t.Fatalf("failed validation: Initialization has succeeded even though there are no arguments: %v", client)
		}
	})
}

func TestGetIssues(t *testing.T) {
	t.Run("when correct response", func(t *testing.T) {
		mux := http.NewServeMux()
		mux.HandleFunc(
			"/issues.json",
			successResponseHandler,
		)
		h := httptest.NewServer(mux)
		defer h.Close()

		client, err := NewClient(
			h.URL,
			"dummyapikey",
			false,
		)

		if err != nil {
			t.Fatalf("failed by NewClient(): %v", err)
		}

		res, _ := client.GetIssues()

		var expect interface{}

		expect = 10
		if got := res.Issues[0].ID; got != expect {
			t.Fatalf("assert failed: ID got = %v, but want: %v", got, expect)
		}

		expect = "Ticket title"
		if got := res.Issues[0].Subject; got != expect {
			t.Fatalf("assert failed: Subject got = %v, but want: %v", got, expect)
		}
	})

	t.Run("when bad response", func(t *testing.T) {
		mux := http.NewServeMux()
		mux.HandleFunc(
			"/issues.json",
			badStatusResponseHandler,
		)
		h := httptest.NewServer(mux)
		defer h.Close()

		client, err := NewClient(
			h.URL,
			"dummyapikey",
			false,
		)

		_, err = client.GetIssues()

		if err == nil {
			t.Fatalf("Should be Error if status code is other than 200.")
		}
	})
}

func TestCreateIssue(t *testing.T) {
	// dummy struct for CreateIssue() execution.
	issue := PostIssueContent{
		ProjectID: 1,
		TrackerID: 1,
		StatusID:  1,
		Subject:   "dummy",
	}

	t.Run("when correct response", func(t *testing.T) {
		mux := http.NewServeMux()
		mux.HandleFunc(
			"/issues.json",
			successCreateIssueResponseHandler,
		)
		h := httptest.NewServer(mux)
		defer h.Close()

		client, _ := NewClient(
			h.URL,
			"dummyapikey",
			false,
		)
		res, _ := client.CreateIssue(issue)

		var expect interface{}

		expect = 11
		if got := res.Issue.ID; got != expect {
			t.Fatalf("assert failed: ID got = %v, but want: %v", got, expect)
		}

		expect = "Ticket title 2"
		if got := res.Issue.Subject; got != expect {
			t.Fatalf("assert failed: Subject got = %v, but want: %v", got, expect)
		}
	})

	t.Run("when bad response", func(t *testing.T) {
		mux := http.NewServeMux()
		mux.HandleFunc(
			"/issues.json",
			badStatusResponseHandler,
		)
		h := httptest.NewServer(mux)
		defer h.Close()

		client, _ := NewClient(
			h.URL,
			"dummyapikey",
			false,
		)
		_, err := client.CreateIssue(issue)

		if err == nil {
			t.Fatalf("Should be Error if status code is other than 201.")
		}
	})
}
