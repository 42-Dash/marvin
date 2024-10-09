package git

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
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
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(
		"Authorization",
		fmt.Sprintf("Bearer %s", os.Getenv("GITHUB_ACCESS")),
	)

	return client.Do(req)
}
