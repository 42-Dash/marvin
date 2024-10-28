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
	if fileExists(config.Repo + "/" + EXECUTABLE_NAME) {
		err := os.Remove(config.Repo + "/" + EXECUTABLE_NAME)
		if err != nil {
			return fmt.Errorf("error removing old executable: %v", err)
		}
	}

	if !fileExists(config.Repo+"/Makefile") && !fileExists(config.Repo+"/makefile") {
		return fmt.Errorf("Makefile not found")
	}

	cmd := exec.Command("/usr/bin/make", "-C", config.Repo)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("%v", stderr.String())
	}

	return nil
}

func selectGradingFunction(league string) func(string, string, int) (int, error) {
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
	if _, err := os.Stat(config.Tracesfile); err == nil {
		os.Remove(config.Tracesfile)
	}

	logger, err := traces.NewLogger(config.Tracesfile)
	if err != nil {
		return err
	}
	defer logger.CloseLogger()

	if err = compileProject(config); err != nil {
		logger.CompilationError(err.Error())
		return nil
	}

	var gradingFunction = selectGradingFunction(config.League)

	for _, repo := range config.Maps {
		res, err := gradingFunction(
			config.Repo+"/"+EXECUTABLE_NAME,
			repo.Path,
			repo.Timeout,
		)
		if err == nil {
			logger.GradingSuccess(repo.Path, res)
		} else if err.Error() == "timeout" {
			logger.TimeoutError(repo.Path)
		} else {
			logger.GradingError(repo.Path, res, err.Error())
		}
	}

	return nil
}
