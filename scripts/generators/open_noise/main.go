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
	rows, cols           int
	W, A, E              int
	range_min, range_max int
	printSolution        bool
	start_row, start_col int
	final_row, final_col int
	filename             string
)

func randomPop(list []rune) ([]rune, rune) {
	index := rand.Intn(len(list))
	poped := list[index]
	list = append(list[:index], list[index+1:]...)
	return list, poped
}

func generateOpenMap(surface []rune) string {
	var content strings.Builder = strings.Builder{}
	var nums []int = []int{0, 0, 0, 0, 0, 0, 0, 0,  0}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if j == start_col && i == start_row {
				content.WriteString("MM")
			} else if j == final_col && i == final_row {
				content.WriteString("GG")
			} else {
				new_surface, poped := randomPop(surface)
				surface = new_surface
				content.WriteByte(byte(poped))

				random := rand.Intn(range_max-range_min+1) + range_min
				nums[random - 1]++
				content.WriteByte(byte('0' + random))
			}
		}
		content.WriteString("\n")
	}
	fmt.Println("Distribution:", nums)
	return content.String()
}

func getSurfaces() []rune {
	map_surf := rows*cols - 2

	w_total := int(float64(W*map_surf) / float64(W+A+E))
	a_total := int(float64(A*map_surf) / float64(W+A+E))
	e_total := map_surf - (w_total + a_total)
	fmt.Println("Square:", map_surf+2)
	fmt.Println("Ration:", w_total, ":", a_total, ":", e_total)

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
	if len(os.Args) != 6 && len(os.Args) != 7 {
		log.Fatal("Usage: ./map_generator [size rows:cols] [proportion W:A:E] [start row:col] [goal row:col] <output_file_name>")
	}

	utils.ParseArr(os.Args[1], ":", &rows, &cols)
	utils.ParseArr(os.Args[2], ":", &W, &A, &E)
	utils.ParseArr(os.Args[3], ":", &range_min, &range_max)
	utils.ParseArr(os.Args[4], ":", &start_row, &start_col)
	utils.ParseArr(os.Args[5], ":", &final_row, &final_col)

	if rows < 1 || cols < 1 || rows*cols < 2 {
		log.Fatal("Rows and cols must be greater than 1")
	}

	if W < 0 || A < 0 || E < 0 {
		log.Fatal("W, A and E must be greater than 0")
	}

	if range_min < 0 || range_max < 0 || range_min > range_max || range_max > 9 {
		log.Fatal("Range min and max must be between 0 and 9")
	}

	if start_row < 0 || start_col < 0 || start_row >= rows || start_col >= cols {
		log.Fatal("Start row and col must be within the map")
	}

	if final_row < 0 || final_col < 0 || final_row >= rows || final_col >= cols {
		log.Fatal("Goal row and col must be within the map")
	}

	if start_row == final_row && start_col == final_col {
		log.Fatal("Start and goal must be different")
	}

	if len(os.Args) == 6 {
		printSolution = true
	} else {
		filename = os.Args[6]
	}
}
