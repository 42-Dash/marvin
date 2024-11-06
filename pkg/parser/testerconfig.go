package parser

import (
	"dashinette/internals/traces"
	"encoding/json"
	"os"
)

const MAPS_FILE string = "config/maps.json"

type Map struct {
	Path    string `json:"path"`
	Name    string `json:"name"`
	Timeout int    `json:"timeout"`
}

type MapsJSON struct {
	RookieMaps []Map `json:"rookieleague"`
	OpenMaps   []Map `json:"openleague"`
}

type TesterConfig struct {
	Maps []Map
	Args TesterArgs
}

type TesterArgs struct {
	TeamName string `json:"teamname"`
	RepoPath string `json:"repo"`
	League   string `json:"league"`
}

func SerializeTesterConfig(team Team, repo, tracesfile string) string {
	config := TesterArgs{
		TeamName: team.Name,
		RepoPath: traces.GetRepoPathContainerized(repo),
		League:   team.League,
	}

	value, _ := json.Marshal(config)
	return string(value)
}

func DeserializeTesterConfig(data []byte) (TesterConfig, error) {
	var args TesterArgs

	err := json.Unmarshal(data, &args)
	if err != nil {
		return TesterConfig{}, err
	}

	var maps MapsJSON
	file, err := os.Open(MAPS_FILE)
	if err != nil {
		return TesterConfig{}, err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&maps)
	if err != nil {
		return TesterConfig{}, err
	}
	var config = TesterConfig{
		Args: args,
	}
	if args.League == "rookie" {
		config.Maps = maps.RookieMaps
	} else {
		config.Maps = maps.OpenMaps
	}
	return config, err
}