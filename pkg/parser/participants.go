package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

const ParticipantsFile = "config/participants.json"

// Team represents a team of participants.
type Team struct {
	Name      string   `json:"name"`
	Nicknames []string `json:"members"`
	League    string   `json:"league"`
}

// The team name can only contain:
//   - ASCII letters
//   - digits
//   - characters '.' '-' and '_'
func (t *Team) Validate() error {
	re := regexp.MustCompile(`^[a-zA-Z0-9\.\-_]+$`)

	if !re.MatchString(t.Name) {
		return fmt.Errorf("invalid team name: %s", t.Name)
	}
	return nil
}

// Participants represents a list of teams.
// Contains the entire content of the participants.json file.
type Participants struct {
	Teams []Team `json:"teams"`
}

// Validates if the team names are unique.
func (p *Participants) validateDuplicates() error {
	var errMessage string
	unique := make(map[string]bool)

	for _, team := range p.Teams {
		if unique[team.Name] {
			if errMessage == "" {
				errMessage = fmt.Sprintf("team name is not unique: %s", team.Name)
			} else {
				errMessage = fmt.Sprintf("%s; team name is not unique: %s", errMessage, team.Name)
			}
		}
		unique[team.Name] = true
	}

	if errMessage != "" {
		return fmt.Errorf("validation errors: %v", errMessage)
	}
	return nil
}

// Validate validates the participants object.
func (p *Participants) validateTeams() error {
	if len(p.Teams) == 0 {
		return fmt.Errorf("no teams found")
	}
	var errs string
	for _, team := range p.Teams {
		if err := team.Validate(); err != nil {
			if len(errs) == 0 {
				errs = err.Error()
			} else {
				errs = fmt.Sprintf("%s; %s", errs, err)
			}
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("validation errors: %v", errs)
	}
	return nil
}

// Validates the participants object.
// Checks if there are any teams and if each team name is valid and unique.
func (p *Participants) Validate() error {
	var errs string

	if err := p.validateTeams(); err != nil {
		errs = err.Error()
	}

	if err := p.validateDuplicates(); err != nil {
		if len(errs) == 0 {
			errs = err.Error()
		} else {
			errs = fmt.Sprintf("%s; %s", errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("validation errors: %v", errs)
	}
	return nil
}

// Loads the participants from the given file.
//
// Parameters:
//   - filename: The name of the file to load the participants from.
//
// Returns:
//   - Participants: The participants object.
//   - error: An error object if an error occurred, otherwise nil.
func LoadParticipantsJSON() (Participants, error) {
	file, err := os.Open(ParticipantsFile)
	if err != nil {
		return Participants{}, err
	}
	defer file.Close()

	participants := Participants{}
	err = json.NewDecoder(file).Decode(&participants)
	if err != nil {
		return Participants{}, err
	}
	if err := participants.Validate(); err != nil {
		return Participants{}, err
	}
	return participants, nil
}
