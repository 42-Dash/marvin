package main

import (
	utils "dashinette/scripts/generators"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

var (
	rows, cols    int
	min, max      int
	printSolution bool
	filename      string
)

func generateRookieMap() string {
	var content strings.Builder = strings.Builder{}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if j == 0 && i == 0 {
				content.WriteRune('M')
			} else if j == cols-1 && i == rows-1 {
				content.WriteRune('G')
			} else {
				random := rand.Intn(max-min+1) + min
				content.WriteByte(byte('0' + random))
			}
		}
		content.WriteString("\n")
	}

	return content.String()
}

func main() {
	content := generateRookieMap()

	if printSolution {
		fmt.Print(content)
	} else {
		err := os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Done!")
	}
}

func init() {
	if len(os.Args) != 3 && len(os.Args) != 4 {
		log.Fatal("Usage: ./map_generator [size rows:cols] [range min:max] <output_file_name>")
	}

	utils.ParseArr(os.Args[1], ":", &rows, &cols)
	utils.ParseArr(os.Args[2], ":", &min, &max)

	if min < 1 || max > 9 || min > max {
		log.Fatal("Min and max must be between 1 and 9 and min must be less than max")
	}

	if rows < 1 || cols < 1 || rows*cols < 2 {
		log.Fatal("Rows and cols must be greater than 1")
	}

	if len(os.Args) == 3 {
		printSolution = true
	} else {
		filename = os.Args[3]
	}
}
