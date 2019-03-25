package out

import (
	"strconv"
	"time"

	"github.com/dobassy/concourse-redmine-resource"
	"github.com/dobassy/concourse-redmine-resource/redmine"
)

// Put is the main logic in the OUT resouce
func Put(request Request) (Response, error) {
	client, err := redmine.NewClient(
		request.Source.URI,
		request.Source.Apikey,
	)
	if err != nil {
		return Response{}, err
	}

	issue, err := client.CreateIssue(
		redmine.PostIssueContent{
			Subject:   request.Params.Subject,
			ProjectID: request.Params.ProjectID,
			TrackerID: request.Params.TrackerID,
			StatusID:  request.Params.StatusID,
		})
	if err != nil {
		return Response{}, err
	}

	return Response{
		Version: resource.Version{
			Timestamp: time.Now(),
		},
		Metadata: []resource.MetadataPair{
			{
				Name:  "ticket id",
				Value: strconv.Itoa(issue.Issue.ID),
			},
			{
				Name:  "project name",
				Value: issue.Issue.Project.Name,
			},
		},
	}, nil
}
