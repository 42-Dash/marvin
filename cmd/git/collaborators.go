package git

import (
	"fmt"
	"net/http"
	"os"
)

// Returns the URL to add a collaborator to a repository.
func addCollaboratorsURL(repo string, nickname string) string {
	return fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/collaborators/%s",
		os.Getenv("GITHUB_ORGANISATION"),
		repo,
		nickname,
	)
}

// Returns an error message based on the status code.
func addCollaboratorsErrorMessage(statusCode int) string {
	switch statusCode {
	case http.StatusForbidden:
		return "forbidden"
	case http.StatusNotFound:
		return "not found"
	default:
		return "unexpected status code"
	}
}

// Adds list of collaborators to a repository.
//
// Parameters:
//   - repo: The name of the repository.
//   - usernames: A list of usernames to add as collaborators.
//
// Returns:
//   - error: An error object if an error occurred, otherwise nil.
func AddCollaborators(repo string, usernames []string) error {
	var errs string

	for _, username := range usernames {
		url := addCollaboratorsURL(repo, username)
		req, err := sendRequest(http.MethodPut, url, nil)
		if err != nil {
			// this structure is repeated in the other functions as well
			// TODO: think about how you can refactor this
			if len(errs) == 0 {
				errs = err.Error()
			} else {
				errs = fmt.Sprintf("%s; %s", errs, err)
			}
		}
		defer req.Body.Close()

		if req.StatusCode != http.StatusCreated && req.StatusCode != http.StatusNoContent {
			errorMessage := fmt.Errorf("cannot add collaborator %s: %s",
				username,
				addCollaboratorsErrorMessage(req.StatusCode),
			)
			if len(errs) == 0 {
				errs = errorMessage.Error()
			} else {
				errs = fmt.Sprintf("%s; %s", errs, errorMessage)
			}
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("error: %v", errs)
	}
	return nil
}
