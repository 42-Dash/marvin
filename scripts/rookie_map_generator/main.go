package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	row, _ := strconv.Atoi(os.Args[1])
	col, _ := strconv.Atoi(os.Args[2])

	content := generateRookieMap(row, col)

	if len(os.Args) == 4 {
		fmt.Print(content)
	} else {
		err := os.WriteFile(os.Args[3], []byte(content), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func init() {
	if len(os.Args) != 3 && len(os.Args) != 4 {
		log.Fatal("Usage: ./map_generator [rows] [cols] <output_file_name>")
	}

	rows, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	cols, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	if rows < 1 || cols < 1 || rows*cols < 2 {
		log.Fatal("Rows and cols must be greater than 1")
	}
}
