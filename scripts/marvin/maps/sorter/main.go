package main

import (
	"bytes"
	"dashinette/internals/traces"
	"dashinette/pkg/parser"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
)

var (
	resultsFile string
	mapsFile    string
	league      string
)

const (
	rookieLeague = "rookieleague"
	openLeague   = "openleague"
)

type Element struct {
	Name  string
	Index int
}

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	}
	return false
}

func parseFile(filename string) traces.Results {
	var content traces.Results

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	err = json.Unmarshal(data, &content)
	if err != nil {
		log.Fatalf("Error parsing file: %v", err)
	}

	return content
}

func sortResults(results *traces.Results, order map[string]Element) {
	var sortedLevels []traces.Level = make([]traces.Level, len(results.Levels))

	for _, level := range results.Levels {
		sortedLevels[order[level.Name].Index] = level
		sortedLevels[order[level.Name].Index].Name = order[level.Name].Name
	}

	results.Levels = sortedLevels
}

func serializeResults(results traces.Results, filename string) {
	original, err := json.Marshal(results)
	if err != nil {
		log.Fatalf("Error serializing results: %v", err)
	}

	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, original, "", "\t")

	err = os.WriteFile(filename, prettyJSON.Bytes(), 0644)
	if err != nil {
		log.Fatalf("Error writing file: %v", err)
	}
}

func getOrder(file string) map[string]Element {
	var content parser.MapsJSON

	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	err = json.Unmarshal(data, &content)
	if err != nil {
		log.Fatalf("Error parsing file: %v", err)
	}

	var maps = []parser.Map{}
	if league == rookieLeague {
		maps = content.RookieMaps
	} else {
		maps = content.OpenMaps
	}

	var order map[string]Element = make(map[string]Element)
	for idx, mapData := range maps {
		order[path.Base(mapData.Path)] = Element{mapData.Name, idx}
	}

	return order
}

func main() {
	var results traces.Results = parseFile(resultsFile)
	var order map[string]Element = getOrder(mapsFile)
	sortResults(&results, order)
	serializeResults(results, resultsFile)
}

func init() {
	fmt.Println("Please make sure you have a backup of the original file.")
	fmt.Scanln()
	if len(os.Args) != 4 {
		log.Fatal("usage: ./main <results file> <maps file with actual order> [league]")
	}

	resultsFile = os.Args[1]
	if !fileExists(resultsFile) {
		log.Fatalf("error: results file %s does not exist", resultsFile)
	}

	mapsFile = os.Args[2]
	if !fileExists(mapsFile) {
		log.Fatalf("error: maps.json file %s does not exist", mapsFile)
	}

	league = os.Args[3]
	if league != rookieLeague && league != openLeague {
		log.Fatalf("error: league must be either rookieleague or openleague")
	}
}
