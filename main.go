package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"teste/cmd/scripts"
)

func main() {
	scripts.LoadParticipantsAndCreateRepos("participants.json")
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	const (
		access = "GITHUB_ACCESS"
		org    = "GITHUB_ORGANISATION"
	)

	for _, env := range []string{access, org} {
		if os.Getenv(env) == "" {
			log.Fatalf("Error: %s not found in .env", env)
		}
	}
}
