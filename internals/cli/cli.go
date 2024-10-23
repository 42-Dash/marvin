package cli

import (
	"dashinette/pkg/parser"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
)

// Constants for the different options in the CLI.
const (
	CreateRepoValue           = "Create repositories"
	AddCollaboratorValue      = "Add collaborator (push access)"
	SetReposReadOnlyValue     = "Set repositories to read-only"
	GradeWorksValue           = "Grade works (create results.json)"
	GradeWorksWithTracesValue = "Grade works with traces (push traces)"
	ExitValue                 = "Exit"
)

// headerTemplate is the template for the header of the CLI.
const headerTemplate = `+---------------------------------------------+
|                    Dash                     |
+---------------------------------------------+
| Welcome to Dash!                            |
|                                             |
| Through this CLI you can:                   |
| - Create repositories in your organization  |
| - Add collaborators to repositories         |
| - Modify collaborator permissions           |
| - Grade works (create results.json)         |
| - Grade works with traces (push traces)     |
|                                             |
+---------------------------------------------+
`

// rerenderHeader clears the terminal and prints the header.
func rerenderHeader(header string) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Print(header)
}

// InteractiveCLI is the main entry point for the CLI.
//
// Parameters:
//   - filename: The name of the file to load the participants from.
//
// The function uses logs to print the status of the operation.
func InteractiveCLI(settings parser.Participants) {
	prompt := promptui.Select{
		Label: "Select an option",
		Items: []string{
			CreateRepoValue,
			AddCollaboratorValue,
			SetReposReadOnlyValue,
			GradeWorksValue,
			GradeWorksWithTracesValue,
			ExitValue,
		},
	}
	rerenderHeader(headerTemplate)

	for {
		_, result, err := prompt.Run()
		if err != nil {
			log.Fatal(err)
		}
		rerenderHeader(headerTemplate)

		switch result {
		case CreateRepoValue:
			createRepos(settings)
		case AddCollaboratorValue:
			addCollaborators(settings)
		case SetReposReadOnlyValue:
			setReposReadOnly(settings)
		case GradeWorksValue:
			gradeWorks(settings)
		case GradeWorksWithTracesValue:
			gradeWorksWithTraces(settings)
		case ExitValue:
			fmt.Println("Goodbye!")
			os.Exit(0)
		}
	}
}
