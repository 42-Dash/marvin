package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	row, col int
}

type Node struct {
	point Point
	g     int
	h     float64
	path  string
}

func (n Node) f() float64 {
	return float64(n.g) + n.h
}

func neighbors(content []string, p Point) []Node {
	var neighbors []Node
	var steps []string = []string{"U", "D", "L", "R"}

	for idx, delta := range []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		nn := Point{p.row + delta.row, p.col + delta.col}

		if nn.row < 0 || nn.row >= len(content) || nn.col < 0 || nn.col >= len(content[0]) {
			continue
		} else {
			neighbors = append(neighbors, Node{nn, 0, 0, steps[idx]})
		}
	}
	return neighbors
}

func searchIndex(open []Node, neighbor Node) int {
	// binary search
	low, high := 0, len(open)-1
	for low <= high {
		mid := (low + high) / 2
		if open[mid].f() < neighbor.f() || (open[mid].f() == neighbor.f() && open[mid].g < neighbor.g) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func Search(content []string) (string, float64) {
	start := findChar(content, 'M')
	end := findChar(content, 'G')

	open := []Node{{start, 0, heuristic(start, end), ""}}
	closed := make(map[Point]bool)

	for len(open) > 0 {
		current := open[0]
		open = open[1:]

		if current.point.col == end.col && current.point.row == end.row {
			return current.path, current.f()
		}

		closed[current.point] = true

		for _, neighbor := range neighbors(content, current.point) {
			if closed[neighbor.point] {
				continue
			}

			if content[neighbor.point.row][neighbor.point.col] >= '1' && content[neighbor.point.row][neighbor.point.col] <= '9' {
				neighbor.g = current.g + int(content[neighbor.point.row][neighbor.point.col] - '0')
			} else {
				neighbor.g = current.g
			}
			neighbor.h = heuristic(neighbor.point, end)
			neighbor.path = current.path + neighbor.path

			insertIndex := searchIndex(open, neighbor)
			open = append(open[:insertIndex], append([]Node{neighbor}, open[insertIndex:]...)...)
		}
	}

	return "", 0
}

var averageStep = 0.0

func heuristic(a, b Point) float64 {
	return (math.Abs(float64(a.row-b.row)) + math.Abs(float64(a.col-b.col))) * averageStep
}

func findChar(content []string, char rune) Point {
	for row, line := range content {
		for col, c := range line {
			if char == c {
				return Point{row, col}
			}
		}
	}
	log.Fatal("Character not found")
	return Point{}
}

func calculateAverageStep(content []string) {
	var total int
	for _, line := range content {
		for _, char := range line {
			if char >= '1' && char <= '9' {
				total += int(char - '0')
			}
		}
	}
	averageStep = float64(total) / float64(len(content)*len(content[0])-2)
}

func readfile(inputfile string) []string {
	file, err := os.ReadFile(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	return slices.DeleteFunc(
		strings.Split(string(file), "\n"),
		func(s string) bool { return s == "" },
	)
}

func main() {
	content := readfile(os.Args[1])

	if len(os.Args) == 3 {
		var err error
		averageStep, err = strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		calculateAverageStep(content)
	}
	path, _ := Search(content)
	fmt.Println(path)
}

func init() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		log.Fatal("Wrong argc")
	}
}
