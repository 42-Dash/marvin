package github

import (
	"fmt"
	"net/http"
	"os"
)

// Returns the URL to delete a repository in the organization.
func deleteRepoUrl(name string) string {
	return fmt.Sprintf(
		"https://api.github.com/repos/%s/%s",
		os.Getenv("GITHUB_ORGANISATION"),
		name,
	)
}

// Deletes a repository in the organization.
//
// Parameters:
//   - name: The name of the repository to delete.
//
// Returns:
//   - error: An error object if an error occurred, otherwise nil.
func DeleteRepo(name string) error {
	url := deleteRepoUrl(name)
	res, err := sendRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	return nil
}
