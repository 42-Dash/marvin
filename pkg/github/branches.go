package github

import "fmt"

// Checks if a branch exists in the repository.
func isBranchExist(repoPath string, branch string) bool { // todo
	err := executeCommand(
		repoPath,
		"git",
		"show-ref",
		"--verify",
		"--quiet",
		fmt.Sprintf("refs/heads/%s", branch),
	)
	return err != nil
}

// Creates a branch in the repository.
func SwitchBranch(repoPath string, branch string) (err error) {

	if isBranchExist(repoPath, branch) {
		err = executeCommand(repoPath, "git", "checkout", "-b", branch)
	} else {
		err = executeCommand(repoPath, "git", "switch", branch)
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
