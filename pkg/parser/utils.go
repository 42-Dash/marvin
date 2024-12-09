package parser

import (
	"dashinette/pkg/constants/marvin"
	"fmt"
	"path/filepath"
	"strings"
)

// Returns the path to the repository of the given team.
func GetRepoPath(name string) string {
	return fmt.Sprintf(marvin.DASH_FOLDER+"repos/%s", name)
}

// Returns the path to the traces file of the given team.
func GetTracesPath(name string) string {
	return fmt.Sprintf(marvin.DASH_FOLDER+"traces/%s.json", name)
}

func GetRepoPathContainerized(path string) string {
	return strings.Replace(
		filepath.ToSlash(path),
		marvin.DASH_FOLDER+"repos/",
		"repo/",
		1,
	)
}

func GetTracesPathContainerized(name string) string {
	return fmt.Sprintf("traces/%s.json", name)
}
