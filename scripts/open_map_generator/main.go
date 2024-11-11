package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func randomPop(list []rune) ([]rune, rune) {
	index := rand.Intn(len(list))
	poped := list[index]
	list = append(list[:index], list[index+1:]...)
	return list, poped
}

func generateOpenMap(rows, cols int, surface []rune) string {
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

func getSurfaces(row, col int) []rune {
	w, _ := strconv.Atoi(os.Args[3])
	a, _ := strconv.Atoi(os.Args[4])
	e, _ := strconv.Atoi(os.Args[5])
	map_surf := row*col - 2

	w_total := int(float64(w*map_surf) / float64(w+a+e))
	a_total := int(float64(a*map_surf) / float64(w+a+e))
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
	row, _ := strconv.Atoi(os.Args[1])
	col, _ := strconv.Atoi(os.Args[2])

	surfaces := getSurfaces(row, col)
	content := generateOpenMap(row, col, surfaces)

	if len(os.Args) == 6 {
		fmt.Print(content)
	} else {
		err := os.WriteFile(os.Args[6], []byte(content), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func init() {
	if len(os.Args) != 6 && len(os.Args) != 7 {
		log.Fatal("Usage: ./map_generator [rows] [cols] [W] [A] [E] <output_file_name>")
	}

	for _, str := range os.Args[1:6] {
		if _, err := strconv.Atoi(str); err != nil {
			log.Fatal(err)
		}
	}

	rows, _ := strconv.Atoi(os.Args[1])
	cols, _ := strconv.Atoi(os.Args[2])

	if rows < 1 || cols < 1 || rows*cols < 2 {
		log.Fatal("Rows and cols must be greater than 1")
	}
}
