package out

import (
	"github.com/dobassy/concourse-redmine-resource"
)

// Request is received by Stdin
type Request struct {
	Source resource.Source `json:"source"`
	Params Params          `json:"params"`
}

type Params struct {
	Subject     string `json:"subject"`
	ProjectID   int    `json:"project_id"`
	TrackerID   int    `json:"tracker_id"`
	StatusID    int    `json:"status_id"`
	ContentFile string `json:"content_file"`
}

type Response struct {
	Version  resource.Version        `json:"version"`
	Metadata []resource.MetadataPair `json:"metadata"`
}
