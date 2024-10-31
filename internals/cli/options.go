package cli

import (
	"dashinette/internals/containerization"
	"dashinette/internals/traces"
	"dashinette/pkg/github"
	"dashinette/pkg/parser"
	"log"
)

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

func pushSubjects(participants parser.Participants) {
	for _, team := range participants.Teams {
		err := github.UploadFileToRoot(traces.GetRepoPath(team.Name), "README.md", "add subjects", "main")
		if err != nil {
			log.Printf("Error pushing subjects for team %s: %v", team.Name, err)
		} else {
			log.Printf("Successfully pushed subjects for team %s", team.Name)
		}
	}
}

func evaluateAssignments(participants parser.Participants) {
	for _, team := range participants.Teams {
		err := github.CloneRepo(team.Name, traces.GetRepoPath(team.Name))
		if err != nil {
			log.Printf("Error cloning repo for team %s: %v", team.Name, err)
		} else {
			log.Printf("Successfully cloned repo for team %s", team.Name)
		}
	}
	for _, team := range participants.Teams {
		err := containerization.GradeAssignmentInContainer(team, traces.GetRepoPath(team.Name), traces.GetTracesPath(team.Name))
		if err != nil {
			log.Printf("Error grading works for team %s: %v", team.Name, err)
		} else {
			log.Printf("Successfully graded works for team %s", team.Name)
		}
	}
}

func pushTraces(participants parser.Participants) {
	for _, team := range participants.Teams {
		err := github.UploadFileToRoot(traces.GetRepoPath(team.Name), traces.GetTracesPath(team.Name), "Upload traces", "traces")
		if err != nil {
			log.Printf("Error pushing results for team %s: %v", team.Name, err)
		} else {
			log.Printf("Successfully pushed results for team %s", team.Name)
		}
	}
}

func createResults(participants parser.Participants) {
	// for _, team := range participants.Teams {
	// 	err := github.CreateBranch(team.Name, "traces")
	// 	if err != nil {
	// 		log.Printf("Error creating branch for team %s: %v", team.Name, err)
	// 	} else {
	// 		log.Printf("Successfully created branch for team %s", team.Name)
	// 	}
	// }
}

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
