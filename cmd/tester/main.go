package main

import (
	"dashinette/internals/grader"
	"dashinette/pkg/parser"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		config, err := parser.DeserializeTesterConfig([]byte(os.Args[1]))
		fmt.Println("Config: ", config)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		grader.MultistageGraderWithTraces(config)
	} else {
		log.Fatalf("Error: wrong number of arguments")
	}
}
