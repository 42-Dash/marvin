package scripts

import (
	"log"
	"teste/cmd/git"
	"teste/cmd/parser"
)

func LoadParticipantsAndDeleteRepos(filename string) {
	participants, err := parser.LoadParticipants(filename)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, team := range participants.Teams {
		err := git.DeleteRepo(team.Name)
		if err != nil {
			log.Printf("Error deleting repo for team %s: %v", team.Name, err)
		} else {
			log.Printf("Successfully deleted repo for team %s", team.Name)
		}
	}
}

func LoadParticipantsAndCreateRepos(filename string) {
	participants, err := parser.LoadParticipants(filename)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, team := range participants.Teams {
		err := git.CreateRepo(team.Name, true)
		if err != nil {
			log.Printf("Error creating repo for team %s: %v", team.Name, err)
		} else {
			log.Printf("Successfully created repo for team %s", team.Name)
		}
	}
}
