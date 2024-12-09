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
	noiseMaps    int    = 6
	resultFolder string = "../../../../../dashes/marvin/maps/rookieleague"
	imageFolder  string = "../../../../../dashes/marvin/maps/images"
)

var (
	imageGenerationScript = []string{"scripts/marvin/maps/generators/rookie_image", "go", "run", "main.go"}
	noiseGenerationScript = []string{"scripts/marvin/maps/generators/rookie_noise", "go", "run", "main.go"}
)

var parametrs = [][]string{
	{"50", "50", "1:9", "0:0", "-1:-1"},
	{"100", "100", "1:9", "-1:-1", "0:0"},
	{"150", "150", "1:9", "0:0", "-1:-1"},
	{"100", "100", "1:3", "-1:0", "0:-1"},
	{"100", "100", "5:9", "0:0", "-1:-1"},
	{"100", "100", "8:9", "9:-10", "-10:9"},
	{"300", "300", imagePath(6), "t", "0:-1", "-1:0"},
	{"250", "250", imagePath(7), "f", "2:-3", "-5:2"},
	{"400", "400", imagePath(8), "f", "-5:5", "5:-5"},
	{"400", "400", imagePath(9), "f", "5:-5", "-5:5"},
	{"400", "600", imagePath(10), "f", "-5:5", "5:-4"},
	{"400", "400", imagePath(11), "f", "25:25", "-25:-25"},
	{"400", "200", imagePath(12), "f", "-5:-5", "20:-80"},
	{"300", "450", imagePath(13), "f", "20:-105", "-5:-20"},
	{"400", "300", imagePath(14), "f", "-23:10", "10:50"},
	{"400", "400", imagePath(15), "t", "30:-30", "-23:170"},
	{"450", "450", imagePath(16), "t", "10:10", "-13:-13"},
	{"500", "500", imagePath(17), "f", "10:10", "-13:-13"},
	{"500", "900", imagePath(18), "f", "10:200", "-13:-11"},
	{"500", "900", imagePath(19), "t", "10:-10", "-43:413"},
	{"500", "900", imagePath(20), "t", "70:110", "-33:-213"},
	{"500", "900", imagePath(21), "t", "-20:15", "-213:313"},
	{"500", "900", imagePath(22), "t", "0:-1", "-1:0"},
}

func executeCommand(dir, command string, args ...string) error {
	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = dir

	return cmd.Run()
}

func imagePath(i int) string {
	return fmt.Sprintf("%v/rookieleague_level_%d.png", imageFolder, i)
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
