package grader

import (
	"bytes"
	"dashinette/internals/grader/open"
	"dashinette/internals/grader/rookie"
	"dashinette/internals/logger"
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
		// logger.Error.Printf("Team %s Makefile not found", config.Args.TeamName)
		return fmt.Errorf("error: Makefile not found")
	}

	cmd := exec.Command("/usr/bin/make", "-C", config.Args.RepoPath)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		// logger.Error.Printf("Team %s error compiling project: %v", config.Args.TeamName, err)
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
	_, err := os.Stat(traces.GetTracesPath(config.Args.TeamName))
	if err == nil {
		os.Remove(traces.GetTracesPath(config.Args.TeamName))
	}
	tr := traces.NewLogger()
	defer tr.StoreInFile(traces.GetTracesPathContainerized(config.Args.TeamName))


	if err := compileProject(config); err != nil {
		tr.AddCompilation(err.Error())
		return nil
	} else {
		tr.AddCompilation("OK")
	}

	var gradingFunction = selectGradingFunction(config.Args.League)

	for _, repo := range config.Maps {
		_, res, err := gradingFunction(
			config.Args.RepoPath+"/"+EXECUTABLE_NAME,
			repo.Path,
			repo.Timeout,
		)
		if err == nil {
			tr.AddStage(repo.Path, res, "OK")
			logger.Info.Printf("Team %s graded map %s: %d", config.Args.TeamName, repo.Path, res)
		} else {
			tr.AddStage(repo.Path, res, err.Error())
			logger.Info.Printf("Team %s error grading map %s: %v", config.Args.TeamName, repo.Path, err)
		}
	}

	return nil
}
