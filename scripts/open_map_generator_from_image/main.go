package main

import (
	"fmt"
	"image/png"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	row, col int
}

type RGB struct {
	r, g, b uint32
}

var (
	size     Point
	start    Point
	finish   Point
	filename string
	inverted bool
)

func readImage() [][]RGB {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	ans := [][]RGB{}
	bounds := img.Bounds()
	width, heights := bounds.Max.X, bounds.Max.Y

	for i := 0; i < size.row; i++ {
		line := []RGB{}
		for j := 0; j < size.col; j++ {
			r, g, b, _ := img.At(j*width/size.col, i*heights/size.row).RGBA()
			line = append(line, RGB{r / 257, g / 257, b / 257})
		}
		ans = append(ans, line)
	}

	return ans
}

func getSurface(rgb RGB) (byte, uint8) {
	var value uint8
	var surf byte

	if rgb.r == rgb.g && rgb.r == rgb.b {
		surf = "AEW"[rand.Int()%3]
	} else if rgb.g > rgb.r && rgb.g > rgb.b {
		surf = 'E'
	} else if rgb.b > rgb.r && rgb.b > rgb.g {
		surf = 'W'
	} else if rgb.r > rgb.g && rgb.r > rgb.b {
		surf = 'A'
	} else if rgb.r == rgb.g {
		surf = "AE"[rand.Int()%2]
	} else if rgb.r == rgb.b {
		surf = "AW"[rand.Int()%2]
	} else {
		surf = "EW"[rand.Int()%2]
	}

	if inverted {
		value = uint8((rgb.r + rgb.g + rgb.b) / 95)
	} else {
		value = 8 - uint8((rgb.r+rgb.g+rgb.b)/95)
	}

	return surf, value
}

func generateOpenMap(surfaces [][]RGB) string {
	var content strings.Builder = strings.Builder{}

	for i := 0; i < len(surfaces); i++ {
		for j := 0; j < len(surfaces[i]); j++ {
			if i == start.row && j == start.col {
				content.WriteString("MM")
			} else if i == finish.row && j == finish.col {
				content.WriteString("GG")
			} else {
				surf, value := getSurface(surfaces[i][j])
				content.WriteByte(surf)
				content.WriteByte(byte('1' + value))
			}
		}
		content.WriteString("\n")
	}

	return content.String()
}

func main() {
	surfaces := readImage()
	content := generateOpenMap(surfaces)

	if len(os.Args) == 6 {
		fmt.Print(content)
	} else {
		err := os.WriteFile(os.Args[6], []byte(content), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func parsePoint(point string) Point {
	parts := strings.Split(point, ":")

	if len(parts) != 2 {
		log.Fatal("Invalid point format, must be [row:col]")
	}

	row, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}

	col, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}

	if col < 0 || row < 0 {
		log.Fatal("Row and col must be greater than 0")
	}

	return Point{row, col}
}

func init() {
	log.SetFlags(0)
	if len(os.Args) != 6 && len(os.Args) != 7 {
		log.Fatal("Usage: ./map_generator [size rows:cols] [start row:col] [end row:col] [image_file.png] [invert option t/f] <output_file_name>")
	}

	size = parsePoint(os.Args[1])
	start = parsePoint(os.Args[2])
	finish = parsePoint(os.Args[3])
	filename = os.Args[4]
	inverted = os.Args[5] == "t"

	if start.row >= size.row || start.col >= size.col {
		log.Fatal("Start point is out of bounds")
	}

	if finish.row >= size.row || finish.col >= size.col {
		log.Fatal("Finish point is out of bounds")
	}

	if start.col == finish.col && start.row == finish.row {
		log.Fatal("Start and finish points are the same")
	}

	if _, err := os.Stat(filename); err != nil {
		log.Fatal(err)
	}

	if os.Args[5] != "t" && os.Args[5] != "f" {
		log.Fatal("Invert option must be t (for true) or f (for false)")
	}

	if size.row < 1 || size.col < 1 || size.row*size.col < 2 {
		log.Fatal("Not enough space for start and finish points")
	}
}
