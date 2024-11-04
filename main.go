package main

import (
	"dashinette/internals/cli"
	"dashinette/internals/logger"
	"dashinette/pkg/parser"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

const participantsFile = "participants.json"

// Main function to start the CLI.
// It prompts the user to enter the path to the participants file.
// Then it loads the participants and starts the interactive CLI.
func main() {
	participants, err := parser.LoadParticipantsJSON(participantsFile)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	logger.InitLogger()
	defer logger.CloseFile()

	cli.InteractiveCLI(participants)
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
		"DOCKER_IMAGE_NAME",
	}

	for _, env := range variables {
		if os.Getenv(env) == "" {
			log.Fatalf("Error: %s not found in .env", env)
		}
	}

	imageName := os.Getenv("DOCKER_IMAGE_NAME")

	// Check if the Docker image exists
	cmd := exec.Command("docker", "image", "inspect", imageName)
	if err := cmd.Run(); err != nil {
		log.Printf("Docker image %s not found. Building it...", imageName)
		buildCmd := exec.Command("docker", "build", "-t", imageName, ".")
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr
		if err := buildCmd.Run(); err != nil {
			log.Fatalf("Failed to build Docker image %s: %v", imageName, err)
		}
	}
}
