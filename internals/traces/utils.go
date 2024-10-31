package traces

import (
	"fmt"
	"path/filepath"
	"strings"
)

const DashFolder = "dashes/marvin/"

// Returns the path to the repository of the given team.
func GetRepoPath(name string) string {
	return fmt.Sprintf(DashFolder+"repos/%s", name)
}

// Returns the path to the traces file of the given team.
func GetTracesPath(name string) string {
	return fmt.Sprintf(DashFolder+"traces/%s.json", name)
}

func GetRepoPathContainerized(path string) string {
	return strings.Replace(
		filepath.ToSlash(path),
		DashFolder+"repos/",
		"repo/",
		1,
	)
}

func GetTracesPathContainerized(name string) string {
	return fmt.Sprintf("traces/%s.json", name)
}
