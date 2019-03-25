package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dobassy/concourse-redmine-resource/in"
)

func main() {
	var request in.Request
	decoder := json.NewDecoder(os.Stdin)
	if err := decoder.Decode(&request); err != nil {
		fatal("failed to decode", err)
	}

	response, err := in.Get(request)
	if err != nil {
		fatal("running Get command", err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		fatal("writing response to stdout", err)
	}
}

func fatal(message string, err error) {
	fmt.Fprintf(os.Stderr, "error %s: %s\n", message, err)
	os.Exit(1)
}
