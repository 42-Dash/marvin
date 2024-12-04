package main

import (
	utils "dashinette/scripts/marvin/maps/generators"
	"fmt"
	"image/png"
	"log"
	"math/rand"
	"os"
	"strings"
)

type RGB struct {
	r, g, b uint32
}

var (
	size          utils.Point
	start         utils.Point
	finish        utils.Point
	imageFile     string
	inverted      bool
	printSolution bool
	filename      string
)

func readImage() [][]RGB {
	file, err := os.Open(imageFile)
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

	for i := 0; i < size.Row; i++ {
		line := []RGB{}
		for j := 0; j < size.Col; j++ {
			r, g, b, _ := img.At(j*width/size.Col, i*heights/size.Row).RGBA()
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
	var surfDist []int = []int{0, 0, 0}
	var nbrDist []int = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < len(surfaces); i++ {
		for j := 0; j < len(surfaces[i]); j++ {
			if i == start.Row && j == start.Col {
				content.WriteString("MM")
			} else if i == finish.Row && j == finish.Col {
				content.WriteString("GG")
			} else {
				surf, value := getSurface(surfaces[i][j])
				content.WriteByte(surf)
				{
					if surf == 'W' {
						surfDist[0]++
					} else if surf == 'A' {
						surfDist[1]++
					} else {
						surfDist[2]++
					}
					nbrDist[value]++
				}
				content.WriteByte(byte('1' + value))
			}
		}
		content.WriteString("\n")
	}
	fmt.Println("Surface distribution:", surfDist)
	fmt.Println("Number distribution:", nbrDist)

	return content.String()
}

func main() {
	surfaces := readImage()
	content := generateOpenMap(surfaces)

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
	log.SetFlags(0)
	if len(os.Args) != 6 && len(os.Args) != 7 {
		log.Fatal("Usage: ./map_generator [size rows:cols] [start row:col] [end row:col] [image_file.png] [invert option t/f] <output_file_name>")
	}

	utils.ParseArr(os.Args[1], ":", &size.Row, &size.Col)
	utils.ParseArr(os.Args[2], ":", &start.Row, &start.Col)
	utils.ParseArr(os.Args[3], ":", &finish.Row, &finish.Col)
	imageFile = os.Args[4]
	inverted = os.Args[5] == "t"

	if start.Row < 0 {
		start.Row = size.Row + start.Row
	}

	if start.Col < 0 {
		start.Col = size.Col + start.Col
	}

	if finish.Row < 0 {
		finish.Row = size.Row + finish.Row
	}

	if finish.Col < 0 {
		finish.Col = size.Col + finish.Col
	}

	if start.Row < 0 || start.Col < 0 || finish.Row < 0 || finish.Col < 0 {
		log.Fatal("Out of bounds")
	}

	if start.Row >= size.Row || start.Col >= size.Col || finish.Row >= size.Row || finish.Col >= size.Col {
		log.Fatal("Out of bounds")
	}

	if start.Col == finish.Col && start.Row == finish.Row {
		log.Fatal("Start and finish points are the same")
	}

	if _, err := os.Stat(imageFile); err != nil {
		log.Fatal(err)
	}

	if os.Args[5] != "t" && os.Args[5] != "f" {
		log.Fatal("Invert option must be t (for true) or f (for false)")
	}

	if size.Row < 1 || size.Col < 1 || size.Row*size.Col < 2 {
		log.Fatal("Not enough space for start and finish points")
	}

	if len(os.Args) == 6 {
		printSolution = true
	} else {
		filename = os.Args[6]
	}
}
