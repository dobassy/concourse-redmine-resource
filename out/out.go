package out

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dobassy/concourse-redmine-resource"
	"github.com/dobassy/concourse-redmine-resource/redmine"
)

// Put is the main logic in the OUT resouce
func Put(client redmine.Client, request Request) (Response, error) {

	postIssueContent, err := buildContent(request)
	if err != nil {
		return Response{}, err
	}

	issue, err := client.CreateIssue(postIssueContent)
	if err != nil {
		return Response{}, err
	}

	return Response{
		Version: resource.Version{
			Timestamp: time.Now(),
		},
		Metadata: []resource.MetadataPair{
			{
				Name:  "Issue ID",
				Value: strconv.Itoa(issue.Issue.ID),
			},
			{
				Name:  "Project Name",
				Value: issue.Issue.Project.Name,
			},
			{
				Name:  "Subject",
				Value: issue.Issue.Subject,
			},
			{
				Name:  "Issue URL",
				Value: fmt.Sprintf("%s/issues/%d", request.Source.URI, issue.Issue.ID),
			},
		},
	}, nil
}

func buildContent(request Request) (redmine.PostIssueContent, error) {
	issueContent := redmine.PostIssueContent{
		Subject:     request.Params.Subject,
		ProjectID:   request.Params.ProjectID,
		TrackerID:   request.Params.TrackerID,
		StatusID:    request.Params.StatusID,
		Description: "",
	}

	// request.Params.Subject is ignored if request.Params.ContentFile is passed
	if request.Params.ContentFile != "" {
		content, err := readContentFile(request.Params.ContentFile)
		fmt.Printf("request.Params.ContentFile: %v", request.Params.ContentFile)
		if err != nil {
			return redmine.PostIssueContent{}, err
		}
		issueContent.Subject = content["subject"]
		issueContent.Description = content["description"]
	}

	return issueContent, nil
}

func readContentFile(path string) (map[string]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return map[string]string{}, err
	}
	defer file.Close()

	content := map[string]string{"subject": "", "description": ""}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			content["subject"] = line
		} else {
			content["description"] += fmt.Sprintln(line)
		}
		i++
	}

	return content, nil
}
