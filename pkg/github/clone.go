package github

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5"
)

// Returns the URL to clone the given repository.
func cloneRepoUrl(repo string) string {
	return fmt.Sprintf("https://github.com/%s/%s",
		os.Getenv("GITHUB_ORGANISATION"),
		repo,
	)
}

// CloneRepo clones the given repository to the given directory.
//
// Parameters:
//   - repo: The name of the repository to clone.
//   - dir: The directory to clone the repository to.
//
// Returns:
//   - An error if the cloning failed.
func CloneRepo(repo string, dir string) error {
	url := cloneRepoUrl(repo)

	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: os.Getenv("GITHUB_ORGANISATION"),
			Password: os.Getenv("GITHUB_ACCESS"),
		},
		URL:      url,
		Progress: os.Stdout,
	})

	fmt.Printf("err: %v\n", err)

	return err
}
