package github

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"
)

// The permission levels for collaborators.
const (
	// The collaborator can push to the repository.
	PUSH = "push"
	// The collaborator can read only from the repository.
	READ = "read"
)

// Creates an authenticated request to the GitHub API.
//
// Parameters:
//   - method: The HTTP method to use.
//   - url: The URL to send the request to.
//   - payload: The payload to send with the request.
//
// Returns:
//   - *http.Response: The response object.
//   - error: An error object if an error occurred, otherwise nil.
func sendRequest(
	method string,
	url string,
	payload []byte,
) (*http.Response, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(
		"Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GITHUB_ACCESS")),
	)

	return client.Do(req)
}

// executeCommand executes a command in the given directory.
func executeCommand(dir, command string, args ...string) error {
	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = dir

	return cmd.Run()
}
