package main

import (
	"dashinette/internals/cli"
	"dashinette/pkg/parser"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// TODO: Think about the best way to handle the participants file.
const participantsFile = "participants.json"

// Main function to start the CLI.
// It prompts the user to enter the path to the participants file.
// Then it loads the participants and starts the interactive CLI.
func main() {
	participants, err := parser.LoadParticipantsJSON(participantsFile)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	cli.InteractiveCLI(participants)

	// a, b := os.Args[1], os.Args[2]
	// c, _ := strconv.Atoi(os.Args[3])

	// should outpur "Hello"
	// grade, err := grading.GradeRookieLeagueAssignment(a, b, c)
	// fmt.Println("Grade: ", grade, "Error: ", err)

	// // should be graded
	// grade, err = grading.GradeAssignment("tests/rookie", "tests/planet.txt", 3);
	// fmt.Println("Grade: ", grade, "Error: ", err);

	// // timeout error
	// grade, err = grading.GradeAssignment("tests/timeout", "planet.txt", 3);
	// fmt.Println("Grade: ", grade, "Error: ", err);

	// // file not found
	// grade, err = grading.GradeAssignment("tests/doesnt exist", "planet.txt", 3);
	// fmt.Println("Grade: ", grade, "Error: ", err);
}

// Checks if all required environment variables are set.
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var variables []string = []string{
		"GITHUB_ACCESS",
		"GITHUB_ORGANISATION",
	}

	for _, env := range variables {
		if os.Getenv(env) == "" {
			log.Fatalf("Error: %s not found in .env", env)
		}
	}
}
