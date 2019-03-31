package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dobassy/concourse-redmine-resource/out"
	"github.com/dobassy/concourse-redmine-resource/redmine"
)

func main() {
	var request out.Request
	decoder := json.NewDecoder(os.Stdin)
	if err := decoder.Decode(&request); err != nil {
		fatal("failed to decode", err)
	}

	if request.Params.ContentFile != "" {
		request.Params.ContentFile = filepath.Join(os.Args[1], request.Params.ContentFile)

		if err := fileExists(request.Params.ContentFile); err == false {
			fatal(
				fmt.Sprintf("content file (%v) not found", request.Params.ContentFile),
				errors.New("no err because it is a bool value"),
			)
		}
	}

	client, err := redmine.NewClient(
		request.Source.URI,
		request.Source.Apikey,
	)
	if err != nil {
		fatal("client initialization", err)
	}

	response, err := out.Put(client, request)
	if err != nil {
		fatal("running command", err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		fatal("writing response to stdout", err)
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func fatal(message string, err error) {
	fmt.Fprintf(os.Stderr, "error %s: %s\n", message, err)
	os.Exit(1)
}
