package github

import (
	"fmt"
	"path/filepath"
)

// Checks if a branch exists in the repository.
func isBranchExist(name string, branch string) bool {
	repoPath := filepath.Join("repo", name)

	err := executeCommand(
		repoPath,
		"git",
		"show-ref",
		"--verify",
		"--quiet",
		fmt.Sprintf("refs/heads/%s", branch),
	)
	return err == nil
}

// Creates a branch in the repository.
func CreateBranch(name string, branch string) (err error) {
	repoPath := filepath.Join("repo", name)

	if isBranchExist(name, branch) {
		err = executeCommand(repoPath, "git", "checkout", "-b", branch)
	} else {
		err = executeCommand(repoPath, "git", "checkout", branch)
	}

	if err != nil {
		return fmt.Errorf("failed to create branch: %w", err)
	}

	err = executeCommand(repoPath, "git", "push", "origin", branch)
	if err != nil {
		return fmt.Errorf("failed to push branch: %w", err)
	}
	return nil
}
