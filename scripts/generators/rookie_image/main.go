package main

import (
	utils "dashinette/scripts/generators"
	"fmt"
	"image/png"
	"log"
	"math"
	"os"
	"strings"
)

var (
	rows, cols    int
	inverted      bool
	imageFile     string
	printSolution bool
	filename      string
)

func readImage() [][]uint8 {
	file, err := os.Open(imageFile)
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
	for i := 0; i < rows; i++ {
		line := []uint8{}
		for j := 0; j < cols; j++ {
			r, g, b, _ := img.At(j*width/cols, i*heights/rows).RGBA()
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
	surfaces := readImage()
	content := generateRookieMap(surfaces)

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
	if len(os.Args) != 4 && len(os.Args) != 5 {
		log.Fatal("Usage: ./map_generator [size rows:cols] [image_file.png] [invert option t/f] <output_file_name>")
	}

	utils.ParseArr(os.Args[1], ":", &rows, &cols)

	imageFile = os.Args[2]
	if _, err := os.Stat(imageFile); err != nil {
		log.Fatal(err)
	}

	if rows < 1 || cols < 1 || rows*cols < 2 {
		log.Fatal("Rows and cols must be greater than 1")
	}

	if os.Args[3] != "t" && os.Args[3] != "f" {
		log.Fatal("Invert option must be t (for true) or f (for false)")
	}

	inverted = os.Args[3] == "t"

	if len(os.Args) == 4 {
		printSolution = true
	} else {
		filename = os.Args[4]
	}
}
