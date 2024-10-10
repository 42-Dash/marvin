package cli

import (
	"dashinette/pkg/github"
	"dashinette/pkg/parser"
	"log"
)

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

// Loads the participants from the given file and creates the repositories.
//
// Parameters:
//   - filename: The name of the file to load the participants from.
//
// The function uses logs to print the status of the operation.
func createRepos(participants parser.Participants) {
	for _, team := range participants.Teams {
		err := github.CreateRepo(team.Name, true)
		if err != nil {
			log.Printf("Error creating repo for team %s: %v", team.Name, err)
			continue
		}
		log.Printf("Successfully created repo for team %s", team.Name)
		err = github.SetCollaborators(team.Name, team.Nicknames, github.PUSH)
		if err != nil {
			log.Printf("Error adding collaborators to team %s: %v", team.Name, err)
		} else {
			log.Printf("Successfully added collaborators to team %s", team.Name)
		}
	}
}

// Loads the participants from the given file and restricts the repositories to read-only.
//
// Parameters:
//   - filename: The name of the file to load the participants from.
//
// The function uses logs to print the status of the operation.
func setReposReadOnly(participants parser.Participants) {
	for _, team := range participants.Teams {
		err := github.SetCollaborators(team.Name, team.Nicknames, github.READ)
		if err != nil {
			log.Printf("Error restricting collaborators for team %s: %v", team.Name, err)
		} else {
			log.Printf("Successfully restricted collaborators for team %s", team.Name)
		}
	}
}
