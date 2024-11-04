package cli

import (
	"dashinette/internals/containerization"
	"dashinette/internals/logger"
	"dashinette/internals/traces"
	"dashinette/pkg/github"
	"dashinette/pkg/parser"
)

func createRepos(participants parser.Participants) {
	for _, team := range participants.Teams {
		err := github.CreateRepo(team.Name, true)
		if err != nil {
			logger.Error.Printf("Error creating repo for team %s: %v", team.Name, err)
		} else {
			logger.Info.Printf("Successfully created repo for team %s", team.Name)
		}
	}
}

func addCollaborators(participants parser.Participants) {
	for _, team := range participants.Teams {
		err := github.SetCollaborators(team.Name, team.Nicknames, github.PUSH)
		if err != nil {
			logger.Error.Printf("Error adding collaborators to team %s: %v", team.Name, err)
		} else {
			logger.Info.Printf("Successfully added collaborators to team %s", team.Name)
		}
	}
}

func cloneRepos(participants parser.Participants) (ok bool) {
	ok = true
	for _, team := range participants.Teams {
		err := github.CloneRepo(team.Name, traces.GetRepoPath(team.Name))
		if err != nil {
			logger.Error.Printf("Error cloning repo for team %s: %v", team.Name, err)
			ok = false
		} else {
			logger.Info.Printf("Successfully cloned repo for team %s", team.Name)
		}
	}
	return
}

func pushSubjects(participants parser.Participants) {
	if !cloneRepos(participants) {
		logger.Error.Println("Error cloning repos, cannot push subjects")
		return
	}
	for _, team := range participants.Teams {
		err := github.UploadFileToRoot(traces.GetRepoPath(team.Name), "README.md", "add subjects", "main")
		if err != nil {
			logger.Error.Printf("Error pushing subjects for team %s: %v", team.Name, err)
		} else {
			logger.Info.Printf("Successfully pushed subjects for team %s", team.Name)
		}
	}
}

func evaluateAssignments(participants parser.Participants) {
	if !cloneRepos(participants) {
		logger.Error.Println("Error cloning repos, cannot push subjects")
		return
	}
	for _, team := range participants.Teams {
		err := containerization.GradeAssignmentInContainer(team, traces.GetRepoPath(team.Name), traces.GetTracesPath(team.Name))
		if err != nil {
			logger.Error.Printf("Error grading works for team %s: %v", team.Name, err)
		} else {
			logger.Info.Printf("Successfully graded works for team %s", team.Name)
		}
	}
}

func pushTraces(participants parser.Participants) {
	for _, team := range participants.Teams {
		err := github.UploadFileToRoot(traces.GetRepoPath(team.Name), traces.GetTracesPath(team.Name), "Upload traces", "traces")
		if err != nil {
			logger.Error.Printf("Error pushing results for team %s: %v", team.Name, err)
		} else {
			logger.Info.Printf("Successfully pushed results for team %s", team.Name)
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
			logger.Error.Printf("Error restricting collaborators for team %s: %v", team.Name, err)
		} else {
			logger.Info.Printf("Successfully restricted collaborators for team %s", team.Name)
		}
	}
}
