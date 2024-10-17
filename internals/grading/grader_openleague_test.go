package grading

import "testing"

var VALID_TEST_MAP_OPEN_LEAGUE = []string{"W4A2E3", "MME5GG", "W7E8A9"}

func TestGetScoreOpenLeague_MissingCharacterPoints(t *testing.T) {
	path := "URRD"
	input := VALID_TEST_MAP_OPEN_LEAGUE
	score, err := getScoreOpenLeague(path, input)
	if !(score == 0 && err != nil) {
		t.Errorf("Error: %v", err)
	}
}

func TestGetScoreOpenLeague_CharacterPointsSumMismatchSmaller(t *testing.T) {
	path := "333URRD"
	input := VALID_TEST_MAP_OPEN_LEAGUE
	score, err := getScoreOpenLeague(path, input)
	if !(score == 0 && err != nil) {
		t.Errorf("Error: %v", err)
	}
}

func TestGetScoreOpenLeague_CharacterPointsSumMismatchGreater(t *testing.T) {
	path := "344URRD"
	input := VALID_TEST_MAP_OPEN_LEAGUE
	score, err := getScoreOpenLeague(path, input)
	if !(score == 0 && err != nil) {
		t.Errorf("Error: %v", err)
	}
}

func TestGetScoreOpenLeague_CharacterPointsGreaterThanFive(t *testing.T) {
	path := "361URRD"
	input := VALID_TEST_MAP_OPEN_LEAGUE
	score, err := getScoreOpenLeague(path, input)
	if !(score == 0 && err != nil) {
		t.Errorf("Error: %v", err)
	}
}

func TestGetScoreOpenLeague_CharacterPointsLessThanTwo(t *testing.T) {
	path := "35URRD"
	input := VALID_TEST_MAP_OPEN_LEAGUE
	score, err := getScoreOpenLeague(path, input)
	if !(score == 0 && err != nil) {
		t.Errorf("Error: %v", err)
	}
}

func TestGetScoreOpenLeague_InvalidPath(t *testing.T) {
	path := "334URR"
	input := VALID_TEST_MAP_OPEN_LEAGUE
	score, err := getScoreOpenLeague(path, input)
	if !(score == 0 && err != nil) {
		t.Errorf("Error: %v", err)
	}
}

func TestGetScoreOpenLeague_ValidPath(t *testing.T) {
	path := "154URRD"
	input := VALID_TEST_MAP_OPEN_LEAGUE
	score, err := getScoreOpenLeague(path, input)
	if !(score == 18.5 && err == nil) {
		t.Errorf("Error: invalid score %v", score)
	}
}

func TestGetScoreOpenLeague_EmpyPath(t *testing.T) {
	path := ""
	input := VALID_TEST_MAP_OPEN_LEAGUE
	score, err := getScoreOpenLeague(path, input)
	if !(score == 0 && err != nil) {
		t.Errorf("Error: %v", err)
	}
}

func TestGetScoreOpenLeague_ExceedingBoundsRight(t *testing.T) {
	path := "334URRDR"
	input := VALID_TEST_MAP_OPEN_LEAGUE
	score, err := getScoreOpenLeague(path, input)
	if !(score == 0 && err != nil) {
		t.Errorf("Error: %v", err)
	}
}

func TestGetScoreOpenLeague_ExceedingBoundsLeft(t *testing.T) {
	path := "334LLRRURRD"
	input := VALID_TEST_MAP_OPEN_LEAGUE
	score, err := getScoreOpenLeague(path, input)
	if !(score == 0 && err != nil) {
		t.Errorf("Error: %v", err)
	}
}

func TestGetScoreOpenLeague_ExceedingBoundsUp(t *testing.T) {
	path := "334UUUDDRRD"
	input := VALID_TEST_MAP_OPEN_LEAGUE
	score, err := getScoreOpenLeague(path, input)
	if !(score == 0 && err != nil) {
		t.Errorf("Error: %v", err)
	}
}

func TestGetScoreOpenLeague_ExceedingBoundsDown(t *testing.T) {
	path := "334DDURRU"
	input := VALID_TEST_MAP_OPEN_LEAGUE
	score, err := getScoreOpenLeague(path, input)
	if !(score == 0 && err != nil) {
		t.Errorf("Error: %v", err)
	}
}

func TestGetScoreOpenLeague_InvalidRune(t *testing.T) {
	path := "334DDURRUP"
	input := VALID_TEST_MAP_OPEN_LEAGUE
	score, err := getScoreOpenLeague(path, input)
	if !(score == 0 && err != nil) {
		t.Errorf("Error: %v", err)
	}
}
