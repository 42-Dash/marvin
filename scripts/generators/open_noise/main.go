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
	W, A, E       int
	printSolution bool
	filename      string
)

func randomPop(list []rune) ([]rune, rune) {
	index := rand.Intn(len(list))
	poped := list[index]
	list = append(list[:index], list[index+1:]...)
	return list, poped
}

func generateOpenMap(surface []rune) string {
	var content strings.Builder = strings.Builder{}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if j == 0 && i == 0 {
				content.WriteString("MM")
			} else if j == cols-1 && i == rows-1 {
				content.WriteString("GG")
			} else {
				new_surface, poped := randomPop(surface)
				surface = new_surface
				content.WriteByte(byte(poped))

				random := rand.Intn(9)
				content.WriteByte(byte('1' + random))
			}
		}
		content.WriteString("\n")
	}

	return content.String()
}

func getSurfaces() []rune {
	map_surf := rows*cols - 2

	w_total := int(float64(W*map_surf) / float64(W+A+E))
	a_total := int(float64(A*map_surf) / float64(W+A+E))
	e_total := map_surf - (w_total + a_total)

	surfaces := []rune{}
	for i := 0; i < w_total; i++ {
		surfaces = append(surfaces, 'W')
	}
	for i := 0; i < a_total; i++ {
		surfaces = append(surfaces, 'A')
	}
	for i := 0; i < e_total; i++ {
		surfaces = append(surfaces, 'E')
	}
	return surfaces
}

func main() {
	surfaces := getSurfaces()
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
	if len(os.Args) != 3 && len(os.Args) != 4 {
		log.Fatal("Usage: ./map_generator [size rows:cols] [proportion W:A:E] <output_file_name>")
	}

	utils.ParseArr(os.Args[1], ":", &rows, &cols)
	utils.ParseArr(os.Args[2], ":", &W, &A, &E)

	if rows < 1 || cols < 1 || rows*cols < 2 {
		log.Fatal("Rows and cols must be greater than 1")
	}

	if len(os.Args) == 3 {
		printSolution = true
	} else {
		filename = os.Args[3]
	}
}
