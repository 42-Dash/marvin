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
	CreateRepoValue          = "Create repositories"
	PushSubjectsValue        = "Push subjects"
	AddCollaboratorValue     = "Add collaborators / write access"
	SetReposReadOnlyValue    = "Set repositories to read-only"
	EvaluateSubmissionsValue = "Evaluate the submissions / create traces"
	PushTracesValue          = "Parse & Push traces to 'traces' branch"
	CreateResultsValue       = "Parse & Create results.json file"
	ExitValue                = "Exit"
)

// headerTemplate is the template for the header of the CLI.
const headerTemplate = `+---------------------------------------------+
|                    Menu                     |
+---------------------------------------------+
| Your ad can be here                         |
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
		Label: "Select an action",
		Items: []string{
			CreateRepoValue,
			PushSubjectsValue,
			AddCollaboratorValue,
			SetReposReadOnlyValue,
			EvaluateSubmissionsValue,
			PushTracesValue,
			CreateResultsValue,
			ExitValue,
		},
		Size: 10,
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
		case PushSubjectsValue:
			pushSubjects(settings)
		case AddCollaboratorValue:
			addCollaborators(settings)
		case SetReposReadOnlyValue:
			setReposReadOnly(settings)
		case EvaluateSubmissionsValue:
			evaluateAssignments(settings)
		case PushTracesValue:
			pushTraces(settings)
		case CreateResultsValue:
			createResults(settings)
		case ExitValue:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Not implemented yet")
		}
	}
}
