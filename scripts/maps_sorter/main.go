package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
)

// structure of the required results.json file
type Group struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Score  int    `json:"score"`
	Path   string `json:"path"`
}

type Level struct {
	Name   string   `json:"lvl"`
	Map    []string `json:"map"`
	Groups []Group  `json:"groups"`
}

type Results struct {
	League string  `json:"league"`
	Levels []Level `json:"levels"`
}

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	}

	return false
}

func validateArguments() {
	if len(os.Args) != 2 {
		log.Fatal("usage: ./main <results file>")
	}
	if !fileExists(os.Args[1]) {
		log.Fatalf("error: file %s does not exist", os.Args[1])
	}
}

func parseFile(filename string) Results {
	var content Results

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

func printOrder(heading string, results Results) {
	fmt.Println(heading)
	for i, level := range results.Levels {
		fmt.Printf("Index %d, Level %s\n", i, level.Name)
	}
	fmt.Println()
}

func promptOrder(levels []Level) []int {
	var order []int
	var index int

	fmt.Println("Enter new order of levels:")
	for i, level := range levels {
		fmt.Printf("Index %d, Level %s\n", i, level.Name)
		fmt.Scan(&index)

		if index < 0 || index >= len(levels) {
			log.Fatalf("Invalid index %d", index)
		}
		if slices.Contains(order, index) {
			log.Fatalf("Index %d already used", index)
		}

		order = append(order, index)
	}
	return order
}

func sortResults(results *Results, order []int) {
	var sortedLevels []Level = make([]Level, len(results.Levels))

	for i, index := range order {
		sortedLevels[index] = results.Levels[i]
	}

	results.Levels = sortedLevels
}

func serializeResults(results Results, filename string) {
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

func main() {
	validateArguments()

	var results Results = parseFile(os.Args[1])

	printOrder("Current order of levels:", results)
	var order []int = promptOrder(results.Levels)

	sortResults(&results, order)
	printOrder("\nOrder of levels after sorting:", results)

	serializeResults(results, os.Args[1])
}
