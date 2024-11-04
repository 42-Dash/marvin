package github

import (
	"dashinette/internals/logger"
	"fmt"
	"os"
	"path/filepath"
)

// It adds the file to the repository, commits it and pushes it to the given branch.
//
// Arguments:
// - repo: the repository name (must be cloned beforehands)
// - filename: the file to upload (relative path form the repository root)
// - commit: the commit message
// - branch: the branch to push to
//
// Returns an error if the file does not exist, or if the commands fail.
func UploadFileToRoot(repo, filename, commit, branch string) error {
	if _, err := os.Stat(repo); os.IsNotExist(err) {
		logger.Warn.Printf("Repository %s does not exist", repo)
		logger.Warn.Printf("Cloning %s...", repo)
		if err := CloneRepo(filepath.Base(repo), repo); err != nil {
			return fmt.Errorf("failed to clone repo: %w", err)
		}
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return fmt.Errorf("file %s does not exist", filename)
	}

	if err := SwitchBranch(repo, branch); err != nil {
		return fmt.Errorf("failed to checkout branch: %w", err)
	}
	basename := filepath.Base(filename)

	if err := executeCommand(".", "cp", filename, filepath.Join(repo, basename)); err != nil {
		return fmt.Errorf("failed to copy file: %w to repo %s", err, repo)
	}

	if err := executeCommand(repo, "git", "add", basename); err != nil {
		return fmt.Errorf("failed to add file: %w to repo %s", err, repo)
	}

	err := executeCommand(repo, "git", "commit", "-m", commit)
	if err != nil && err.Error() != "exit status 1" {
		return fmt.Errorf("failed to commit file: %w to repo %s", err, repo)
	}

	if err := executeCommand(repo, "git", "push", "-f", "origin", branch); err != nil {
		return fmt.Errorf("failed to push file: %w to repo %s", err, repo)
	}

	return nil
}
