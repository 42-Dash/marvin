package github

import (
	"fmt"
	"os"
)

// Returns the URL to clone the given repository.
func cloneRepoUrl(repo_name string) string {
	return fmt.Sprintf("https://%s@github.com/%s/%s.git",
		os.Getenv("GITHUB_ACCESS"),
		os.Getenv("GITHUB_ORGANISATION"),
		repo_name,
	)
}

// CloneRepo clones the given repository to the given directory.
//
// Parameters:
//   - repo_name: The name of the repository to clone.
//   - dir: The directory to clone the repository to.
//
// Returns:
//   - An error if the cloning failed.
func CloneRepo(repo_name, destination string) error {
	url := cloneRepoUrl(repo_name)

	if _, err := os.Stat(destination); err == nil {
		if err := os.RemoveAll(destination); err != nil {
			return fmt.Errorf("failed to delete repository: %v", err)
		}
	}

	err := executeCommand(".", "git", "clone", url, destination)

	return err
}
