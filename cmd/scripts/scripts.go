package scripts

import (
	"log"
	"teste/cmd/git"
	"teste/cmd/parser"
)

// Loads the participants from the given file and deletes the repositories.
//
// Parameters:
//   - filename: The name of the file to load the participants from.
//
// The function uses logs to print the status of the operation.
func LoadParticipantsAndDeleteRepos(filename string) {
	participants, err := parser.LoadParticipantsJSON(filename)
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

// Loads the participants from the given file and creates the repositories.
//
// Parameters:
//   - filename: The name of the file to load the participants from.
//
// The function uses logs to print the status of the operation.
func LoadParticipantsAndCreateRepos(filename string) {
	participants, err := parser.LoadParticipantsJSON(filename)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, team := range participants.Teams {
		err := git.CreateRepo(team.Name, true)
		if err != nil {
			log.Printf("Error creating repo for team %s: %v", team.Name, err)
			continue
		}
		log.Printf("Successfully created repo for team %s", team.Name)
		err = git.AddCollaborators(team.Name, team.Nicknames)
		if err != nil {
			log.Printf("Error adding collaborators to team %s: %v", team.Name, err)
		} else {
			log.Printf("Successfully added collaborators to team %s", team.Name)
		}
	}
}
