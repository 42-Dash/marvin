package parser

import (
	"encoding/json"
	"os"
)

const MAPS_FILE string = "maps.json"

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
	Maps       []Map
	Repo       string
	Tracesfile string
	League     string
}

type TesterArgs struct {
	TracesPath string `json:"tracesfile"`
	RepoPath   string `json:"repo"`
	League     string `json:"league"`
}

func SerializeTesterConfig(team Team, repo, tracesfile string) string {
	config := TesterArgs{
		TracesPath: tracesfile,
		RepoPath:   repo,
		League:     team.League,
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
	var config TesterConfig
	config.Repo = args.RepoPath
	config.Tracesfile = args.TracesPath
	config.League = args.League
	if args.League == "rookie" {
		config.Maps = maps.RookieMaps
	} else {
		config.Maps = maps.OpenMaps
	}
	return config, err
}
