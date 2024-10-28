package grader

import (
	"bytes"
	"dashinette/internals/grader/open"
	"dashinette/internals/grader/rookie"
	"dashinette/internals/traces"
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
	if fileExists(config.Args.RepoPath + "/" + EXECUTABLE_NAME) {
		os.Remove(config.Args.RepoPath + "/" + EXECUTABLE_NAME)
	}

	if !fileExists(config.Args.RepoPath + "/Makefile") {
		return fmt.Errorf("error: Makefile not found")
	}

	cmd := exec.Command("/usr/bin/make", "-C", config.Args.RepoPath)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("error: %v", stderr.String())
	}

	return nil
}

func selectGradingFunction(league string) func(string, string, int) (string, int, error) {
	switch league {
	case "rookie":
		return rookie.GradeRookieLeagueAssignment
	case "open":
		return open.GradeOpenLeagueAssignment
	default:
		return nil
	}
}

func MultistageGraderWithTraces(config parser.TesterConfig) error {
	if _, err := os.Stat("traces/" + config.Args.TeamName + ".json"); err == nil {
		os.Remove("traces/" + config.Args.TeamName + ".json")
	}

	logger := traces.NewLogger()
	defer logger.StoreInFile("traces/" + config.Args.TeamName + ".json")

	if err := compileProject(config); err != nil {
		logger.AddCompilation(err.Error())
		return nil
	} else {
		logger.AddCompilation("OK")
	}

	var gradingFunction = selectGradingFunction(config.Args.League)

	for _, repo := range config.Maps {
		_, res, err := gradingFunction(
			config.Args.RepoPath+"/"+EXECUTABLE_NAME,
			repo.Path,
			repo.Timeout,
		)
		if err == nil {
			logger.AddStage(repo.Path, res, "OK")
		} else {
			logger.AddStage(repo.Path, res, err.Error())
		}
	}

	return nil
}
