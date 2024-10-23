package github

import (
	"dashinette/pkg/parser"
	"fmt"
	"os"
)

// pushes the results to the repository.
//
// Parameters:
//   - team: The team to push the results to.
//   - filename: The name of the file to push.
//
// Returns:
//   - error: An error object if an error occurred, otherwise nil.
func PushResults(team parser.Team, filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return fmt.Errorf("file %s does not exist", filename)
	}

	fmt.Printf("Pushing %v to %v repository", filename, team.Name)
	return nil
}
