package cli

import (
	"dashinette/internals/grading"
	"dashinette/pkg/github"
	"dashinette/pkg/parser"
	"fmt"
	"log"
)

// Creates the repositories, without adding any collaborators.
//
// Parameters:
//   - participants: The participants to add as collaborators.
//
// The function uses logs to print the status of the operation.
func createRepos(participants parser.Participants) {
	for _, team := range participants.Teams {
		err := github.CreateRepo(team.Name, true)
		if err != nil {
			log.Printf("Error creating repo for team %s: %v", team.Name, err)
		} else {
			log.Printf("Successfully created repo for team %s", team.Name)
		}
	}
}

// Adds the collaborators to their respective repositories.
//
// Parameters:
//   - participants: The participants to add as collaborators.
//
// The function uses logs to print the status of the operation.
func addCollaborators(participants parser.Participants) {
	for _, team := range participants.Teams {
		err := github.SetCollaborators(team.Name, team.Nicknames, github.PUSH)
		if err != nil {
			log.Printf("Error adding collaborators to team %s: %v", team.Name, err)
		} else {
			log.Printf("Successfully added collaborators to team %s", team.Name)
		}
	}
}

// Grades the works, creating the results.json file.
//
// Parameters:
//   - participants: The participants to add as collaborators.
//
// The function uses logs to print the status of the operation.
func gradeWorks(parser.Participants) {
	// for _, team := range participants.Teams {
	// 	err := github.GradeWorks(team.Name)
	// 	if err != nil {
	// 		log.Printf("Error grading works for team %s: %v", team.Name, err)
	// 	} else {
	// 		log.Printf("Successfully graded works for team %s", team.Name)
	// 	}
	// }
	fmt.Println("Grade works")
}

// utils function that creates the traces file name.
func getTracesFile(team parser.Team) string {
	return fmt.Sprintf("traces/%s.json", team.Name)
}

func getCloningPath(team parser.Team) string {
	return fmt.Sprintf("repo/%s", team.Name)
}

// Grades the works with traces, pushing the results to the repositories.
//
// Parameters:
//   - participants: The participants to add as collaborators.
//
// The function uses logs to print the status of the operation.
func gradeWorksWithTraces(participants parser.Participants) {
	for _, team := range participants.Teams {
		err := github.CloneRepo(team.Name, getCloningPath(team))
		if err != nil {
			log.Printf("Error cloning repo for team %s: %v", team.Name, err)
		} else {
			log.Printf("Successfully cloned repo for team %s", team.Name)
		}
	}
	for _, team := range participants.Teams {
		err := grading.ContainerizedGrader(team, getCloningPath(team), getTracesFile(team))
		if err != nil {
			log.Printf("Error grading works for team %s: %v", team.Name, err)
		} else {
			log.Printf("Successfully graded works for team %s", team.Name)
		}
	}
	for _, team := range participants.Teams {
		err := github.PushResults(team, getTracesFile(team))
		if err != nil {
			log.Printf("Error pushing results for team %s: %v", team.Name, err)
		} else {
			log.Printf("Successfully pushed results for team %s", team.Name)
		}
	}
}

// Restricts the repositories to read-only. (End of the competition)
//
// Parameters:
//   - participants: The participants to add as collaborators.
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
