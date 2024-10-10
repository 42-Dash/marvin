package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"teste/cmd/scripts"
)

const participantsFile = "participants.json"

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s [create|delete|restrict]", os.Args[0])
	}
	switch os.Args[1] {
	case "create":
		scripts.CreateRepos(participantsFile)
	case "delete":
		scripts.DeleteRepos(participantsFile)
	case "restrict":
		scripts.SetReposReadOnly(participantsFile)
	default:
		log.Fatalf("Usage: %s [create|delete]", os.Args[0])
	}
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
