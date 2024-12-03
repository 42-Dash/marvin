package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

var (
	scenario                    int
	rookie_noise_generator      string = "scripts/marvin/maps/generators/rookie_noise/"
	rookie_image_generator      string = "scripts/marvin/maps/generators/rookie_image/"
	rookie_final_map_folder_rel string = "../../../../../dashes/marvin/maps/final_rookie/"
	rookie_final_img_folder_rel string = "../../../../../dashes/marvin/maps/final_rookie_images/"
)

func executeCommand(dir, command string, args ...string) error {
	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = dir

	return cmd.Run()
}

func selectScenario() [][]string {
	var scripts [][]string
	switch scenario {
	case 0:
		scripts = [][]string{
			{"30:30", "1:9", "0:0", "29:29"},
			{"60:60", "1:9", "59:59", "0:0"},
			{"90:90", "1:9", "0:0", "89:89"},
			{"60:60", "1:3", "59:0", "0:59"},
			{"60:60", "5:9", "0:0", "59:59"},
			{"60:60", "8:9", "9:49", "49:9"},
			{"180:180", rookie_final_img_folder_rel + "6.png", "t", "0:179", "179:0"},
		}
	case 1:
		scripts = [][]string{
			{"40:40", "1:9", "0:0", "39:39"},
			{"80:80", "1:9", "79:79", "0:0"},
			{"120:120", "1:9", "0:0", "119:119"},
			{"80:80", "1:3", "79:0", "0:79"},
			{"80:80", "5:9", "0:0", "79:79"},
			{"80:80", "8:9", "9:69", "69:9"},
			{"240:240", rookie_final_img_folder_rel + "6.png", "t", "0:239", "239:0"},
		}
	default:
		scripts = [][]string{
			{"50:50", "1:9", "0:0", "49:49"},
			{"100:100", "1:9", "99:99", "0:0"},
			{"150:150", "1:9", "0:0", "149:149"},
			{"100:100", "1:3", "99:0", "0:99"},
			{"100:100", "5:9", "0:0", "99:99"},
			{"100:100", "8:9", "9:89", "89:9"},
			{"300:300", rookie_final_img_folder_rel + "6.png", "t", "0:299", "299:0"},
		}
	}
	for i := range scripts {
		if i < 6 {
			scripts[i] = append([]string{rookie_noise_generator, "go", "run", "main.go"}, scripts[i]...)
			scripts[i] = append(scripts[i], fmt.Sprintf("%s%d.txt", rookie_final_map_folder_rel, i))
		} else {
			scripts[i] = append([]string{rookie_image_generator, "go", "run", "main.go"}, scripts[i]...)
			scripts[i] = append(scripts[i], fmt.Sprintf("%s%d.txt", rookie_final_map_folder_rel, i))
		}
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
		log.Fatal("usage: ./main [scenario 0-3]")
	}

	var err error
	scenario, err = strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("scenario must be an integer")
	}

	if scenario < 0 || scenario > 3 {
		log.Fatal("scenario must be between 0 and 3")
	}
}
