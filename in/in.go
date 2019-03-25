package in

import (
	"time"

	"github.com/dobassy/concourse-redmine-resource"
)

// Get is the main logic in the IN resoucel.
// Since this resource is OUT (Put) processing only, IN processing
// only sets the version.
func Get(request Request) (Response, error) {
	timestamp := request.Version.Timestamp
	if timestamp.IsZero() {
		timestamp = time.Now()
	}

	return Response{
		Version: resource.Version{
			Timestamp: timestamp,
		},
	}, nil
}
