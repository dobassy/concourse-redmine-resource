package out

import (
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
		Subject:   "subject in test",
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

	t.Run("when valid request", func(t *testing.T) {
		res, err := Put(client, basicRequest)
		if err != nil {
			t.Fatalf("failed to Put(): %+v", err)
		}

		if res.Version.Timestamp.IsZero() {
			t.Fatalf("Timestamp may not be correct: %+v", res)
		}
	})

	t.Run("when invalid request", func(t *testing.T) {
		invalidRequest := basicRequest
		invalidRequest.Params.ContentFile = "nonexistent_file_name"
		_, err := Put(client, invalidRequest)
		if err == nil {
			t.Fatal("error handling may have failed: it must be an error if a non-existent file is specified")
		}
	})
}

func TestBuildContent(t *testing.T) {
	t.Run("when subject is defined but contnet_file is NOT defined", func(t *testing.T) {
		res, err := buildContent(basicRequest)
		if err != nil {
			t.Fatalf("failed to buildContent(): %+v", err)
		}

		var actual interface{}
		var expected interface{}

		actual = res.Subject
		expected = "subject in test"
		if actual != expected {
			t.Fatalf("no match subject: got \"%v\": expected \"%v\"", actual, expected)
		}

		actual = len(res.Description)
		expected = 0
		if actual != expected {
			t.Fatalf("no match description length: got \"%v\": expected \"%v\"", actual, expected)
		}

	})

	t.Run("when only content_file is defined (subject is NOT defined)", func(t *testing.T) {
		request := basicRequest
		// dir := filepath.Dir(os.Args[0])
		request.Params.ContentFile = "fixtures/ticket_content_file.txt"
		request.Params.Subject = ""

		res, err := buildContent(request)
		if err != nil {
			t.Fatalf("failed to buildContent(): %+v", err)
		}

		var actual interface{}
		var expected interface{}

		actual = res.Subject
		expected = "First line is the subject text"
		if actual != expected {
			t.Fatalf("no match subject: got \"%v\": expected \"%v\"", actual, expected)
		}

		actual = len(res.Description)
		expected = 135 // number of characters in fixtures/ticket_content_file.txt
		if actual != expected {
			t.Fatalf("no match description length: got \"%v\": expected \"%v\"", actual, expected)
		}

	})
}

func TestReadContentFile(t *testing.T) {
	t.Run("when test file exists", func(t *testing.T) {
		_, err := readContentFile("fixtures/ticket_content_file.txt")
		if err != nil {
			t.Fatalf("failed to read ContentFile: error %v", err)
		}
	})

	t.Run("when test file NOT exists", func(t *testing.T) {
		_, err := readContentFile("fixtures/not_exists.txt")
		if err == nil {
			t.Fatalf("Error handling can not be performed even though non-existent file is specified.")
		}
	})
}
