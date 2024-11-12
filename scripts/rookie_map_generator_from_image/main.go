package main

import (
	"fmt"
	"image/png"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readImage(row, col int, filename string, inverted bool) [][]uint8 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	ans := [][]uint8{}
	bounds := img.Bounds()
	width, heights := bounds.Max.X, bounds.Max.Y
	for i := 0; i < row; i++ {
		line := []uint8{}
		for j := 0; j < col; j++ {
			r, g, b, _ := img.At(j*width/col, i*heights/row).RGBA()
			r, g, b = r/257, g/257, b/257
			rg := float64(r+g+b) / 96.0
			if inverted {
				line = append(line, uint8(math.Round(rg)))
			} else {
				line = append(line, 8-uint8(math.Round(rg)))
			}
		}
		ans = append(ans, line)
	}

	return ans
}

func generateRookieMap(surfaces [][]uint8) string {
	var content strings.Builder = strings.Builder{}

	for i := 0; i < len(surfaces); i++ {
		for j := 0; j < len(surfaces[i]); j++ {
			if j == 0 && i == 0 {
				content.WriteRune('M')
			} else if j == len(surfaces[i])-1 && i == len(surfaces)-1 {
				content.WriteRune('G')
			} else {
				content.WriteByte(byte('1' + surfaces[i][j]))
			}
		}
		content.WriteString("\n")
	}

	return content.String()
}

func main() {
	row, _ := strconv.Atoi(os.Args[1])
	col, _ := strconv.Atoi(os.Args[2])
	filename := os.Args[3]
	inverted := os.Args[4] == "t"

	surfaces := readImage(row, col, filename, inverted)
	content := generateRookieMap(surfaces)

	if len(os.Args) == 5 {
		fmt.Print(content)
	} else {
		err := os.WriteFile(os.Args[5], []byte(content), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func init() {
	if len(os.Args) != 5 && len(os.Args) != 6 {
		log.Fatal("Usage: ./map_generator [rows] [cols] [image_file.png] [invert option t/f] <output_file_name>")
	}

	rows, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	cols, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(os.Args[3]); err != nil {
		log.Fatal(err)
	}

	if os.Args[4] != "t" && os.Args[4] != "f" {
		log.Fatal("Invert option must be t (for true) or f (for false)")
	}

	if rows < 1 || cols < 1 || rows*cols < 2 {
		log.Fatal("Rows and cols must be greater than 1")
	}
}
