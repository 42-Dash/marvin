package git

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Returns the URL to create a new repository in the organization.
func createRepoUrl() string {
	return fmt.Sprintf(
		"https://api.github.com/orgs/%s/repos",
		os.Getenv("GITHUB_ORGANISATION"),
	)
}

// Returns an error message based on the status code.
func createRepoErrorMessage(statusCode int) string {
	switch statusCode {
	case http.StatusForbidden:
		return "forbidden"
	case http.StatusUnprocessableEntity:
		return "repository already exists"
	default:
		return "unexpected status code"
	}
}

// CreateRepo creates a new repository in the organization.
//
// Parameters:
//   - name: The name of the repository.
//   - is_private: A boolean indicating whether the repository should be private.
//
// Returns:
//   - error: An error object if an error occurred, otherwise nil.
func CreateRepo(name string, is_private bool) error {
	payload, err := json.Marshal(map[string]interface{}{
		"name":    name,
		"private": is_private,
	})
	if err != nil {
		return err
	}

	var url string = createRepoUrl()
	res, err := sendRequest(http.MethodPost, url, payload)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("cannot create repository %s: %s",
			name,
			createRepoErrorMessage(res.StatusCode),
		)
	}
	return nil
}
