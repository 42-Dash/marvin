package main

import (
	"dashinette/internals/cli"
	"dashinette/internals/grading"
	"dashinette/pkg/parser"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// {"tracesfile":"traces/The-Avengers.json","repo":"repo/The-Avengers","league":"rookie"}

// TODO: Think about the best way to handle the participants file.
const participantsFile = "participants.json"

// Main function to start the CLI.
// It prompts the user to enter the path to the participants file.
// Then it loads the participants and starts the interactive CLI.
func main() {
	if len(os.Args) > 1 {
		config, err := parser.DeserializeTesterConfig([]byte(os.Args[1]))
		fmt.Println("Config: ", config)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		grading.MultistageGrader(config)
	} else {
		init_env()
		participants, err := parser.LoadParticipantsJSON(participantsFile)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		cli.InteractiveCLI(participants)
	}
}

// Checks if all required environment variables are set.
func init_env() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var variables []string = []string{
		"GITHUB_ACCESS",
		"GITHUB_ORGANISATION",
		"DOCKER_IMAGE_NAME",
	}

	for _, env := range variables {
		if os.Getenv(env) == "" {
			log.Fatalf("Error: %s not found in .env", env)
		}
	}
}
