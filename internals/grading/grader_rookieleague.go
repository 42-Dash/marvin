package grading

import (
	"fmt"
	"os"
	"strings"
)

// applies the list of given instructions and returns the score.
func getScoreRookieLeague(path string, input []string) (int, error) {
	var x, y, score int

	x, y = playersPosition(input)
	for _, chr := range path {
		if chr == 'U' {
			x -= 1
		} else if chr == 'D' {
			x += 1
		} else if chr == 'L' {
			y -= 1
		} else if chr == 'R' {
			y += 1
		} else {
			return 0, fmt.Errorf("invalid path")
		}

		if x < 0 || x >= len(input) || y < 0 || y >= len(input[0]) {
			return 0, fmt.Errorf("out of bounds")
		}

		if strings.ContainsRune("123456789", rune(input[x][y])) {
			score += int(input[x][y] - '0')
		}
	}
	if input[x][y] != 'G' {
		return 0, fmt.Errorf("marvin didnt reach the goal")
	}
	return score, nil
}

// grades the assignment in the given file.
//
// Parameters:
//   - filename: The name of the file to grade.
//   - timeout: The timeout for the grading process.
//
// Returns:
//   - int: The grade of the assignment.
//   - error: An error object if an error occurred, otherwise nil.
func GradeRookieLeagueAssignment(filename string, inputfile string, timeout int) (int, error) {
	// check if the file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return 0, fmt.Errorf("file not found")
	}

	output, err := executeWithTimeout(filename, inputfile, timeout)
	if err != nil {
		return 0, err
	}

	path, err := extractLastAnswer(output, VALID_RUNES_ROOKIE_LEAGUE)
	if err != nil {
		return 0, err
	}

	input, err := os.ReadFile(inputfile)
	if err != nil {
		return 0, err
	}

	inputStr := strings.Split(string(input), "\n")
	inputStr = inputStr[:len(inputStr)-1]
	score, err := getScoreRookieLeague(path, inputStr)

	if err != nil {
		return 0, err
	}
	return score, nil
}
