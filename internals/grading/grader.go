package grading

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

const (
	// valid runes for the path
	VALID_RUNES_OPEN_LEAGUE   = "012345UDLR"
	VALID_RUNES_ROOKIE_LEAGUE = "UDLR"
)

// returns player position
func playersPosition(input []string) (int, int) {
	for idx, line := range input {
		if strings.Contains(line, "M") {
			return idx, strings.IndexRune(line, 'M')
		}
	}
	return -1, -1
}

// executes the given file with the given input and timeout.
func executeWithTimeout(filename string, input string, timeout int) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, filename, input)

	// Capture stdout
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	// Start the command
	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("failed to start: %v", err)
	}

	// Wait for the command to complete or be killed after 5 seconds
	if err := cmd.Wait(); err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return stdout.String(), nil
		} else {
			return "", fmt.Errorf("unexpected behavior")
		}
	}

	return stdout.String(), nil
}

// returns the last valid answer from the output.
// the last valid answer ends with the last newline rune in the string.
// It consists of valid_runes only.
func extractLastAnswer(output string, valid_runes string) (string, error) {
	var path string

	for _, c := range path {
		if !strings.ContainsRune(valid_runes, c) {
			return "", fmt.Errorf("invalid character in path")
		}
	}

	end := strings.LastIndex(output, "\n")
	if end == -1 {
		return "", fmt.Errorf("no new line found")
	}

	begin := strings.LastIndex(output[:end], "\n")
	if begin == -1 {
		path = output[:end]
	} else {
		path = output[begin+1 : end]
	}

	if len(path) == 0 {
		return "", fmt.Errorf("empty path")
	}
	return path, nil
}
