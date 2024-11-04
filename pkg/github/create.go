package github

import (
	"encoding/base64"
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

func initialCommitUrl(repo_name string) string {
	return fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/contents/README.md",
		os.Getenv("GITHUB_ORGANISATION"),
		repo_name,
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

// initializes the repository creation process with a commit
func initialCommit(repo_name string) error {
	payload, _ := json.Marshal(map[string]interface{}{
		"message": "Initial commit",
		"content": base64.StdEncoding.EncodeToString([]byte("")),
		"branch":  "main",
	})

	url := initialCommitUrl(repo_name)
	res, err := sendRequest(http.MethodPut, url, payload)
	if err != nil {
		return fmt.Errorf("cannot create initial commit: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("cannot create initial commit: %s", err)
	}

	return nil
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

	err = initialCommit(name)
	if err != nil {
		return err
	}

	return nil
}
