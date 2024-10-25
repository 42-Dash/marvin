package grader

import (
	"bytes"
	"dashinette/internals/grader/open"
	"dashinette/internals/grader/rookie"
	"dashinette/pkg/parser"
	"fmt"
	"os"
	"os/exec"
)

const EXECUTABLE_NAME = "marvin"

// Returns true if the file exists.
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Compiles the project and
func compileProject(config parser.TesterConfig) error {
	if fileExists(config.Repo + "/" + EXECUTABLE_NAME) {
		err := os.Remove(config.Repo + "/" + EXECUTABLE_NAME)
		if err != nil {
			return err
		}
	}

	if !fileExists(config.Repo+"/Makefile") && !fileExists(config.Repo+"/makefile") {
		return fmt.Errorf("Makefile not found")
	}

	cmd := exec.Command("/usr/bin/make", "-C", config.Repo)
	// read stdout and stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("Compilation error: %v", stderr.String())
	}

	fmt.Printf("repo {%v}\n", config.Repo)

	return err
}

func MultistageGraderWithTraces(config parser.TesterConfig) (err error) {
	if _, err = os.Stat(config.Tracesfile); err == nil {
		os.Remove(config.Tracesfile)
	}

	traces, err := os.Create(config.Tracesfile)
	if err != nil {
		return
	}
	defer traces.Close()

	if err = compileProject(config); err != nil {
		traces.Write([]byte(fmt.Sprintf("Compilation error: %v", err)))
		return
	}

	var gradingFunction func(string, string, int) (int, error)
	switch config.League {
	case "rookie":
		gradingFunction = rookie.GradeRookieLeagueAssignment
	case "open":
		gradingFunction = open.GradeOpenLeagueAssignment
	}

	for repo := range config.Maps {
		res, err := gradingFunction(config.Repo+"/"+EXECUTABLE_NAME, config.Maps[repo].Path, config.Maps[repo].Timeout)
		if err != nil {
			traces.Write([]byte(fmt.Sprintf("Error grading assignment: %v", err)))
			break
		} else {
			traces.Write([]byte(fmt.Sprintf("Map %v passed with score %v", config.Maps[repo].Name, res)))
		}
	}

	return
}
