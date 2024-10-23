package grading

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

const (
	// Test programs for rookieleague
	ROOKIELEAGUE_NEW_LINE      = "/rookieleague/invalidnewline"
	ROOKIELEAGUE_TIMEOUT       = "/rookieleague/invalidtimeout"
	ROOKIELEAGUE_RUNES         = "/rookieleague/invalidrunes"
	ROOKIELEAGUE_PATH          = "/rookieleague/invalidpath"
	ROOKIELEAGUE_VALID_TIMEOUT = "/rookieleague/validtimeout"
	ROOKIELEAGUE_TEST          = "/rookieleague/validtest"

	// Test programs for openleague
	OPENLEAGUE_MISSED_CHARACTER = "/openleague/invalidmissedcharacter"
	OPENLEAGUE_UNDERFLOW        = "/openleague/invalidunderflow"
	OPENLEAGUE_EXCEEDS_MAP      = "/openleague/invalidexceed"
	OPENLEAGUE_OVERFLOW         = "/openleague/invalidoverflow"
	OPENLEAGUE_NEW_LINE         = "/openleague/invalidnewline"
	OPENLEAGUE_TIMEOUT          = "/openleague/invalidtimeout"
	OPENLEAGUE_RUNES            = "/openleague/invalidrunes"
	OPENLEAGUE_PATH             = "/openleague/invalidpath"
	OPENLEAGUE_NOT_LAST_PATH    = "/openleague/validnotlast"
	OPENLEAGUE_VALID_TIMEOUT    = "/openleague/validtimeout"
	OPENLEAGUE_TEST             = "/openleague/validtest"
)

// slices with values for the tests
var tests = []string{
	ROOKIELEAGUE_NEW_LINE,
	ROOKIELEAGUE_TIMEOUT,
	ROOKIELEAGUE_RUNES,
	ROOKIELEAGUE_PATH,
	ROOKIELEAGUE_VALID_TIMEOUT,
	ROOKIELEAGUE_TEST,
	OPENLEAGUE_MISSED_CHARACTER,
	OPENLEAGUE_UNDERFLOW,
	OPENLEAGUE_EXCEEDS_MAP,
	OPENLEAGUE_OVERFLOW,
	OPENLEAGUE_NEW_LINE,
	OPENLEAGUE_TIMEOUT,
	OPENLEAGUE_RUNES,
	OPENLEAGUE_PATH,
	OPENLEAGUE_NOT_LAST_PATH,
	OPENLEAGUE_VALID_TIMEOUT,
	OPENLEAGUE_TEST,
}

// Constants for the tests
const (
	ROOKIE_LEAGUE_VALID_MAP = "../../testdata/rookieleague/4x4.txt"
	OPEN_LEAGUE_VALID_MAP   = "../../testdata/openleague/4x4.txt"
	BIN                     = "../../bin/"
	CMD                     = "../../cmd/tests/"
	TIMEOUT                 = 3 // seconds
)

func cleanUp() {
	os.RemoveAll(BIN)
}

func TestMain(m *testing.M) {
	// Compile the executables
	for _, test := range tests {
		err := exec.Command("go", "build", "-o", BIN+test, CMD+test).Run()
		// If the compilation fails, remove the executables and exit
		if err != nil {
			cleanUp()
			fmt.Println("go", "build", "-o", BIN+test, CMD+test)
			os.Exit(1)
		}
	}

	// Run the tests
	code := m.Run()

	cleanUp()

	// Exit with the test code
	os.Exit(code)
}

func TestGradeRookieLeagueAssignment_RL_MissedNewLine(t *testing.T) {
	output, err := GradeRookieLeagueAssignment(
		BIN+ROOKIELEAGUE_NEW_LINE,
		ROOKIE_LEAGUE_VALID_MAP,
		TIMEOUT)
	if !(output == 0 && err != nil) {
		t.Errorf("Error: expected output=0 and error, got %v, %v", output, err)
	}
}

func TestGradeRookieLeagueAssignment_RL_InvalidTimeout(t *testing.T) {
	output, err := GradeRookieLeagueAssignment(
		BIN+ROOKIELEAGUE_TIMEOUT,
		ROOKIE_LEAGUE_VALID_MAP,
		TIMEOUT)
	if !(output == 0 && err != nil) {
		t.Errorf("Error: expected output=0 and error, got %v, %v", output, err)
	}
}

func TestGradeRookieLeagueAssignment_RL_InvalidRunes(t *testing.T) {
	output, err := GradeRookieLeagueAssignment(
		BIN+ROOKIELEAGUE_RUNES,
		ROOKIE_LEAGUE_VALID_MAP,
		TIMEOUT)
	if !(output == 0 && err != nil) {
		t.Errorf("Error: expected output=0 and error, got %v, %v", output, err)
	}
}

func TestGradeRookieLeagueAssignment_RL_InvalidPath(t *testing.T) {
	output, err := GradeRookieLeagueAssignment(
		BIN+ROOKIELEAGUE_PATH,
		ROOKIE_LEAGUE_VALID_MAP,
		TIMEOUT)
	if !(output == 0 && err != nil) {
		t.Errorf("Error: expected output=0 and error, got %v, %v", output, err)
	}
}

func TestGradeRookieLeagueAssignment_RL_ValidTimeout(t *testing.T) {
	output, err := GradeRookieLeagueAssignment(
		BIN+ROOKIELEAGUE_VALID_TIMEOUT,
		ROOKIE_LEAGUE_VALID_MAP,
		TIMEOUT)
	if !(output == 8 && err == nil) {
		t.Errorf("Error: expected no error and grade, got error=%v and grade=%v", err, output)
	}
}

func TestGradeRookieLeagueAssignment_RL_ValidTest(t *testing.T) {
	output, err := GradeRookieLeagueAssignment(
		BIN+ROOKIELEAGUE_TEST,
		ROOKIE_LEAGUE_VALID_MAP,
		TIMEOUT)
	if !(output == 8 && err == nil) {
		t.Errorf("Error: expected no error and grade, got error=%v and grade=%v", err, output)
	}
}

func TestGradeOpenLeagueAssignment_OL_MissedCharacter(t *testing.T) {
	output, err := GradeOpenLeagueAssignment(
		BIN+OPENLEAGUE_MISSED_CHARACTER,
		OPEN_LEAGUE_VALID_MAP,
		TIMEOUT)
	if !(output == 0 && err != nil) {
		t.Errorf("Error: expected output=0 and error, got %v, %v", output, err)
	}
}

func TestGradeOpenLeagueAssignment_OL_Underflow(t *testing.T) {
	output, err := GradeOpenLeagueAssignment(
		BIN+OPENLEAGUE_UNDERFLOW,
		OPEN_LEAGUE_VALID_MAP,
		TIMEOUT)
	if !(output == 0 && err != nil) {
		t.Errorf("Error: expected output=0 and error, got %v, %v", output, err)
	}
}

func TestGradeOpenLeagueAssignment_OL_ExceedsMap(t *testing.T) {
	output, err := GradeOpenLeagueAssignment(
		BIN+OPENLEAGUE_EXCEEDS_MAP,
		OPEN_LEAGUE_VALID_MAP,
		TIMEOUT)
	if !(output == 0 && err != nil) {
		t.Errorf("Error: expected output=0 and error, got %v, %v", output, err)
	}
}

func TestGradeOpenLeagueAssignment_OL_Overflow(t *testing.T) {
	output, err := GradeOpenLeagueAssignment(
		BIN+OPENLEAGUE_OVERFLOW,
		OPEN_LEAGUE_VALID_MAP,
		TIMEOUT)
	if !(output == 0 && err != nil) {
		t.Errorf("Error: expected output=0 and error, got %v, %v", output, err)
	}
}

func TestGradeOpenLeagueAssignment_OL_MissedNewLine(t *testing.T) {
	output, err := GradeOpenLeagueAssignment(
		BIN+OPENLEAGUE_NEW_LINE,
		OPEN_LEAGUE_VALID_MAP,
		TIMEOUT)
	if !(output == 0 && err != nil) {
		t.Errorf("Error: expected output=0 and error, got %v, %v", output, err)
	}
}

func TestGradeOpenLeagueAssignment_OL_InvalidTimeout(t *testing.T) {
	output, err := GradeOpenLeagueAssignment(
		BIN+OPENLEAGUE_TIMEOUT,
		OPEN_LEAGUE_VALID_MAP,
		TIMEOUT)
	if !(output == 0 && err != nil) {
		t.Errorf("Error: expected output=0 and error, got %v, %v", output, err)
	}
}

func TestGradeOpenLeagueAssignment_OL_InvalidRunes(t *testing.T) {
	output, err := GradeOpenLeagueAssignment(
		BIN+OPENLEAGUE_RUNES,
		OPEN_LEAGUE_VALID_MAP,
		TIMEOUT)
	if !(output == 0 && err != nil) {
		t.Errorf("Error: expected output=0 and error, got %v, %v", output, err)
	}
}

func TestGradeOpenLeagueAssignment_OL_InvalidPath(t *testing.T) {
	output, err := GradeOpenLeagueAssignment(
		BIN+OPENLEAGUE_PATH,
		OPEN_LEAGUE_VALID_MAP,
		TIMEOUT)
	if !(output == 0 && err != nil) {
		t.Errorf("Error: expected output=0 and error, got %v, %v", output, err)
	}
}

func TestGradeOpenLeagueAssignment_OL_NotLastPath(t *testing.T) {
	output, err := GradeOpenLeagueAssignment(
		BIN+OPENLEAGUE_NOT_LAST_PATH,
		OPEN_LEAGUE_VALID_MAP,
		TIMEOUT)
	if !(output == 13 && err == nil) {
		t.Errorf("Error: expected no error and grade, got error=%v and grade=%v", err, output)
	}
}

func TestGradeOpenLeagueAssignment_OL_ValidTimeout(t *testing.T) {
	output, err := GradeOpenLeagueAssignment(
		BIN+OPENLEAGUE_VALID_TIMEOUT,
		OPEN_LEAGUE_VALID_MAP,
		TIMEOUT)
	if !(output == 13 && err == nil) {
		t.Errorf("Error: expected no error and grade, got error=%v and grade=%v", err, output)
	}
}

func TestGradeOpenLeagueAssignment_OL_ValidTest(t *testing.T) {
	output, err := GradeOpenLeagueAssignment(
		BIN+OPENLEAGUE_TEST,
		OPEN_LEAGUE_VALID_MAP,
		TIMEOUT)
	if !(output == 13 && err == nil) {
		t.Errorf("Error: expected no error and grade, got error=%v and grade=%v", err, output)
	}
}
