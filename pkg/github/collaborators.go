package github

import (
	"fmt"
	"net/http"
	"os"
)

// Returns the URL endpoint to perform crud over collaborators.
func setCollaboratorsURL(repo string, nickname string) string {
	return fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/collaborators/%s",
		os.Getenv("GITHUB_ORGANISATION"),
		repo,
		nickname,
	)
}

// Returns an error message based on the status code.
func setCollaboratorsErrorMessage(statusCode int) string {
	switch statusCode {
	case http.StatusForbidden:
		return "forbidden"
	case http.StatusNotFound:
		return "not found"
	case http.StatusUnprocessableEntity:
		return "unprocessable entity"
	default:
		return "unexpected status code"
	}
}

// Adds list of collaborators to a repository.
//
// Parameters:
//   - repo: The name of the repository.
//   - usernames: A list of usernames to add as collaborators.
//   - pemission: The permission level for the collaborators. (PUSH or READ in this case)
//
// Returns:
//   - error: An error object if an error occurred, otherwise nil.
func SetCollaborators(
	repo string,
	usernames []string,
	pemission string,
) error {
	payload := []byte(fmt.Sprintf("{\"permission\":\"%s\"}", pemission))
	var errs string

	for _, username := range usernames {
		url := setCollaboratorsURL(repo, username)
		req, err := sendRequest(http.MethodPut, url, payload)
		if err != nil {
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
				setCollaboratorsErrorMessage(req.StatusCode),
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
