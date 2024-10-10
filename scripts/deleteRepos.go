package main

import (
	"dashinette/pkg/github"
	"dashinette/pkg/parser"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	participants, err := parser.LoadParticipantsJSON("participants.json")
	if err != nil {
		log.Fatalf("Error loading participants: %v", err)
	}
	deleteRepos(participants)
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

// Loads the participants from the given file and deletes the repositories.
//
// Parameters:
//   - filename: The name of the file to load the participants from.
//
// The function uses logs to print the status of the operation.
func deleteRepos(participants parser.Participants) {
	for _, team := range participants.Teams {
		err := github.DeleteRepo(team.Name)
		if err != nil {
			log.Printf("Error deleting repo for team %s: %v", team.Name, err)
		} else {
			log.Printf("Successfully deleted repo for team %s", team.Name)
		}
	}
}
