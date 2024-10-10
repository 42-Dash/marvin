package scripts

import (
	"log"
	"teste/cmd/git"
	"teste/cmd/parser"
	"github.com/schollz/progressbar/v3"
)

// Loads the participants from the given file and deletes the repositories.
//
// Parameters:
//   - filename: The name of the file to load the participants from.
//
// The function uses logs to print the status of the operation.
func DeleteRepos(participants parser.Participants) {
	bar := progressbar.Default(int64(len(participants.Teams)))

	for _, team := range participants.Teams {
		err := git.DeleteRepo(team.Name)
		if err != nil {
			log.Printf("Error deleting repo for team %s: %v", team.Name, err)
		} else {
			log.Printf("Successfully deleted repo for team %s", team.Name)
		}
		bar.Add(1)
	}
}

// Loads the participants from the given file and creates the repositories.
//
// Parameters:
//   - filename: The name of the file to load the participants from.
//
// The function uses logs to print the status of the operation.
func CreateRepos(participants parser.Participants) {
	bar := progressbar.Default(int64(len(participants.Teams)))

	for _, team := range participants.Teams {
		err := git.CreateRepo(team.Name, true)
		if err != nil {
			log.Printf("Error creating repo for team %s: %v", team.Name, err)
			continue
		}
		log.Printf("Successfully created repo for team %s", team.Name)
		err = git.SetCollaborators(team.Name, team.Nicknames, git.PUSH)
		if err != nil {
			log.Printf("Error adding collaborators to team %s: %v", team.Name, err)
		} else {
			log.Printf("Successfully added collaborators to team %s", team.Name)
		}
		bar.Add(1)
	}
}

// Loads the participants from the given file and restricts the repositories to read-only.
//
// Parameters:
//   - filename: The name of the file to load the participants from.
//
// The function uses logs to print the status of the operation.
func SetReposReadOnly(participants parser.Participants) {
	bar := progressbar.Default(int64(len(participants.Teams)))

	for _, team := range participants.Teams {
		err := git.SetCollaborators(team.Name, team.Nicknames, git.READ)
		if err != nil {
			log.Printf("Error restricting collaborators for team %s: %v", team.Name, err)
		} else {
			log.Printf("Successfully restricted collaborators for team %s", team.Name)
		}
		bar.Add(1)
	}
}
