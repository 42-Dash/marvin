package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// Generates a rookie map with the given number of rows and columns.
func generateRookieMap(rows, cols int) string {
	var content strings.Builder = strings.Builder{}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if j == 0 && i == 0 {
				content.WriteRune('M')
			} else if j == cols-1 && i == rows-1 {
				content.WriteRune('G')
			} else {
				random := rand.Intn(9)
				content.WriteByte(byte('1' + random))
			}
		}
		content.WriteString("\n")
	}

	return content.String()
}

// Generates an open map with the given number of rows and columns.
func generateOpenMap(rows, cols int) string {
	var content strings.Builder = strings.Builder{}
	SURFACES := "WEA"

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if j == 0 && i == 0 {
				content.WriteString("MM")
			} else if j == cols-1 && i == rows-1 {
				content.WriteString("GG")
			} else {
				surf := rand.Intn(3)
				content.WriteByte(SURFACES[surf])
				random := rand.Intn(9)
				content.WriteByte(byte('1' + random))
			}
		}
		content.WriteString("\n")
	}

	return content.String()
}

// Main function to generate the map.
func main() {
	row, _ := strconv.Atoi(os.Args[2])
	col, _ := strconv.Atoi(os.Args[3])
	outputFileName := os.Args[4]

	content := ""
	if os.Args[1] == "open" {
		content = generateOpenMap(row, col)
	} else {
		content = generateRookieMap(row, col)
	}
	err := os.WriteFile(outputFileName, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// Checks if the arguments are valid.
// The first argument must be either 'rookie' or 'open'.
// The second and third arguments must be integers greater than 1.
// The fourth argument must be the name of the output file.
func init() {
	if len(os.Args) != 5 {
		log.Fatal("Usage: map_generator [rookie/open] <rows> <cols> <output_file_name>")
	}

	if os.Args[1] != "rookie" && os.Args[1] != "open" {
		log.Fatal("First argument must be 'rookie' or 'open'")
	}

	rows, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	cols, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	if rows < 2 || cols < 2 {
		log.Fatal("Rows and cols must be greater than 0")
	}
}
