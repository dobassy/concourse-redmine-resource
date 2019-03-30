package redmine

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"
)

var (
	userAgent = fmt.Sprintf("GoClient/%s", runtime.Version())
)

func NewClient(redmineurl string, apikey string) (Client, error) {
	if len(apikey) == 0 {
		return nil, errors.New("missing user apikey")
	}

	parsedURL, err := url.ParseRequestURI(redmineurl)
	if err != nil {
		return nil, errors.New("failed to parse url (e.g. https://your.redmine.fqdn/)")
	}

	return &client{
		apikey: apikey,
		URL:    parsedURL,
		HTTPClient: &http.Client{
			Timeout: 15 * time.Second,
		}}, nil
}

func (c *client) newRequest(method string, endpoint string, body io.Reader) (*http.Request, error) {
	v := url.Values{}
	v.Add("key", c.apikey)

	u := *c.URL
	u.Path = endpoint
	u.RawQuery = v.Encode()

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)

	return req, nil
}

func (c *client) GetIssues() (*IssuesResponse, error) {
	req, err := c.newRequest("GET", "issues.json", nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("http status is not 200: status %v", res.StatusCode))
	}

	var issues IssuesResponse
	if err := decodeBody(res, &issues); err != nil {
		return nil, err
	}

	return &issues, nil
}

func (c *client) CreateIssue(issue PostIssueContent) (*PostIssueResponse, error) {
	issueRequest := PostIssueRequest{
		Issue: issue,
	}
	s, err := json.Marshal(issueRequest)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest("POST", "issues.json", strings.NewReader(string(s)))
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 201 {
		return nil, fmt.Errorf("http status is not 201: status %v", res.StatusCode)
	}

	var issueResponse PostIssueResponse
	if err := decodeBody(res, &issueResponse); err != nil {
		return nil, err
	}

	return &issueResponse, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
