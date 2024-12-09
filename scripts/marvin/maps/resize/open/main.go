package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

var (
	scenario     int
	noiseMaps    int    = 8
	resultFolder string = "../../../../../dashes/marvin/maps/openleague"
	imageFolder  string = "../../../../../dashes/marvin/maps/images"
)

var (
	imageGenerationScript = []string{"scripts/marvin/maps/generators/open_image", "go", "run", "main.go"}
	noiseGenerationScript = []string{"scripts/marvin/maps/generators/open_noise", "go", "run", "main.go"}
)

var parametrs = [][]string{
	{"50", "50", "1:1:1", "1:9", "-3:-1", "1:3"},
	{"80", "80", "2:1:1", "1:9", "3:-1", "-1:3"},
	{"100", "100", "1:1:9", "1:9", "-3:1", "1:-3"},
	{"150", "150", "1:20:20", "1:9", "3:-6", "-6:3"},
	{"200", "200", "1:1:1", "6:9", "-3:1", "1:-3"},
	{"200", "200", "1:1:1", "1:4", "1:-3", "-3:1"},
	{"300", "300", "2:5:6", "6:9", "13:33", "-13:-1"},
	{"300", "300", "5:6:2", "1:4", "-1:33", "3:-3"},
	{"200", "400", "20:30", "-23:-42", imagePath(8), "t"},
	{"300", "300", "20:-30", "-23:42", imagePath(9), "t"},
	{"400", "250", "-20:-30", "23:42", imagePath(10), "f"},
	{"200", "400", "-20:30", "23:-42", imagePath(11), "f"},
	{"300", "600", "20:30", "-23:-42", imagePath(12), "f"},
	{"300", "600", "20:-30", "-23:72", imagePath(13), "f"},
	{"300", "600", "-20:-30", "53:122", imagePath(14), "f"},
	{"300", "600", "-20:40", "44:-142", imagePath(15), "f"},
	{"300", "600", "20:-30", "-23:42", imagePath(16), "f"},
	{"400", "600", "-12:32", "23:-2", imagePath(17), "f"},
	{"380", "600", "20:30", "-23:-42", imagePath(18), "t"},
	{"400", "660", "200:330", "13:12", imagePath(19), "t"},
	{"500", "860", "12:-121", "-13:22", imagePath(20), "t"},
	{"500", "900", "12:-121", "93:22", imagePath(21), "t"},
	{"500", "900", "1:-1", "-1:1", imagePath(22), "t"},
}

func executeCommand(dir, command string, args ...string) error {
	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = dir

	return cmd.Run()
}

func imagePath(i int) string {
	return fmt.Sprintf("%v/openleague_level_%d.png", imageFolder, i)
}

func getDestination(i int) string {
	return fmt.Sprintf("%s/level_%d.txt", resultFolder, i)
}

func selectScenario() [][]string {
	var scripts [][]string

	for _, param := range parametrs {
		rows, _ := strconv.Atoi(param[0])
		cols, _ := strconv.Atoi(param[1])
		rows = rows * scenario / 100
		size := []string{fmt.Sprintf("%d:%d", rows, cols)}
		scripts = append(scripts, append(size, param[2:]...))
	}
	for i := range scripts {
		if i < noiseMaps {
			scripts[i] = append(noiseGenerationScript, scripts[i]...)
		} else {
			scripts[i] = append(imageGenerationScript, scripts[i]...)
		}
		scripts[i] = append(scripts[i], getDestination(i))
	}
	return scripts
}

func main() {
	scenario := selectScenario()
	for _, command := range scenario {
		if err := executeCommand(command[0], command[1], command[2:]...); err != nil {
			fmt.Printf("failed to execute command: %e", err)
			os.Exit(1)
		}
	}
}

func init() {
	fmt.Println("Resize maps")
	if len(os.Args) != 2 {
		log.Fatal("usage: ./main [scenario 0/1/2/3]")
	}

	var err error
	scenario, err = strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("scenario must be an integer")
	}

	if scenario < 0 || scenario > 3 {
		log.Fatal("scenario must be between 0 and 3")
	}

	switch scenario {
	case 0:
		scenario = 60
	case 1:
		scenario = 80
	case 2:
		scenario = 100
	case 3:
		scenario = 120
	}
}
