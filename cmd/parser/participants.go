package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

type Team struct {
	Name      string   `json:"name"`
	Nicknames []string `json:"members"`
}

type Participants struct {
	Teams []Team `json:"teams"`
}

// The repository name can only contain ASCII letters, digits, and the characters ., -, and _.
func (p *Participants) validate() error {
	re := regexp.MustCompile(`^[a-zA-Z0-9\.\-_]+$`)

	for _, team := range p.Teams {
		if !re.MatchString(team.Name) {
			return fmt.Errorf("invalid team name: %s", team.Name)
		}
	}
	return nil
}

func LoadParticipants(filename string) (Participants, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Participants{}, err
	}
	defer file.Close()

	participants := Participants{}
	err = json.NewDecoder(file).Decode(&participants)
	if err != nil {
		return Participants{}, err
	}
	if err := participants.validate(); err != nil {
		return Participants{}, err
	}
	return participants, nil
}
