package cli

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"teste/cmd/parser"
	"teste/cmd/scripts"

	"github.com/manifoldco/promptui"
)

// initTemplate is the template for the initial prompt header.
const initTemplate = `+---------------------------------------------+
| Please enter the path to your               |
| participants.json file to get started.      |
+---------------------------------------------+
`

// PromptParticipantsFile prompts the user to enter the path to the participants file.
//
// Returns:
//   - string: The path to the participants file.
func PromptParticipantsFile() string {
	rerenderHeader(initTemplate)

	prompt := promptui.Prompt{
		Label: "> ",
		Validate: func(input string) error {
			if len(input) < 5 || input[len(input)-5:] != ".json" {
				return fmt.Errorf("invalid file format: %s", input)
			}
			if _, err := os.Stat(input); os.IsNotExist(err) {
				return fmt.Errorf("file not found: %s", input)
			}
			return nil
		},
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	return result
}

// Constants for the different options in the CLI.
const (
	CreateRepoValue = "Create a repository"
	DeleteRepoValue = "Delete a repository"
	SetReposReadOnlyValue = "Set repositories to read-only"
	ExitValue = "Exit"
)

// headerTemplate is the template for the header of the CLI.
const headerTemplate = `+---------------------------------------------+
|                    Dash                     |
+---------------------------------------------+
| Welcome to Dash!                            |
|                                             |
| Through this CLI you can:                   |
| - Create repositories in your organization  |
| - Delete repositories in your organization  |
| - Modify collaborator permissions           |
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
			DeleteRepoValue,
			SetReposReadOnlyValue,
			ExitValue,
		},
	}

	for {
		rerenderHeader(headerTemplate)

		_, result, err := prompt.Run()
		if err != nil {
			log.Fatal(err)
		}

		switch result {
		case CreateRepoValue:
			scripts.CreateRepos(settings)
		case DeleteRepoValue:
			scripts.DeleteRepos(settings)
		case SetReposReadOnlyValue:
			scripts.SetReposReadOnly(settings)
		case ExitValue:
			fmt.Println("Goodbye!")
			os.Exit(0)
		}
	}
}
