package grading

import "testing"

var TEST_VALID_MAP_ROOKIE_LEAGUE = []string{"M9G", "141", "111"}

func TestGetScoreRookieLeague_EmptyPath(t *testing.T) {
	path := ""
	input := TEST_VALID_MAP_ROOKIE_LEAGUE
	score, err := getScoreRookieLeague(path, input)
	if !(score == 0 && err != nil) {
		t.Errorf("Error: %v", err)
	}
}

func TestGetScoreRookieLeague_InvalidPath(t *testing.T) {
	path := "DURL"
	input := TEST_VALID_MAP_ROOKIE_LEAGUE
	score, err := getScoreRookieLeague(path, input)
	if !(score == 0 && err != nil) {
		t.Errorf("Error: %v", err)
	}
}

func TestGetScoreRookieLeague_OutOfBounds(t *testing.T) {
	path := "U"
	input := TEST_VALID_MAP_ROOKIE_LEAGUE
	score, err := getScoreRookieLeague(path, input)
	if !(score == 0 && err != nil) {
		t.Errorf("Error: %v", err)
	}
}

func TestGetScoreRookieLeague_MarvinDidntReachGoal(t *testing.T) {
	path := "R"
	input := TEST_VALID_MAP_ROOKIE_LEAGUE
	score, err := getScoreRookieLeague(path, input)
	if !(score == 0 && err != nil) {
		t.Errorf("Error: %v", err)
	}
}

func TestGetScoreRookieLeague_ValidPath(t *testing.T) {
	path := "DRRU"
	input := TEST_VALID_MAP_ROOKIE_LEAGUE
	score, err := getScoreRookieLeague(path, input)
	if !(score == 6 && err == nil) {
		t.Errorf("Error: %v", err)
	}
}

func TestGetScoreRookieLeague_InvalidRuneNewLine(t *testing.T) {
	path := "DRRU\n"
	input := TEST_VALID_MAP_ROOKIE_LEAGUE
	score, err := getScoreRookieLeague(path, input)
	if !(score == 0 && err != nil) {
		t.Errorf("Error: %v", err)
	}
}

func TestGetScoreRookieLeague_InvalidRune(t *testing.T) {
	path := "DRRUW"
	input := TEST_VALID_MAP_ROOKIE_LEAGUE
	score, err := getScoreRookieLeague(path, input)
	if !(score == 0 && err != nil) {
		t.Errorf("Error: %v", err)
	}
}
