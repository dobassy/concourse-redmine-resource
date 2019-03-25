package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dobassy/concourse-redmine-resource/out"
)

func main() {
	var request out.Request
	decoder := json.NewDecoder(os.Stdin)
	if err := decoder.Decode(&request); err != nil {
		fatal("failed to decode", err)
	}

	//srcpath := os.Args[1]
	response, err := out.Put(request)
	if err != nil {
		fatal("running command", err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		fatal("writing response to stdout", err)
	}
}

func fatal(message string, err error) {
	fmt.Fprintf(os.Stderr, "error %s: %s\n", message, err)
	os.Exit(1)
}
