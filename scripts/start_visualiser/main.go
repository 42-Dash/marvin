package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const DOCKERFILE_PATH = "dashes/marvin/visualiser/"
const DOCKER_IMAGE_NAME_PREFIX = "dashinette-visualiser"

func createDockerImageName(name string) string {
	name = strings.Split(name, ".")[0]
	return DOCKER_IMAGE_NAME_PREFIX + "-" + name
}

func createDockerImage(name string) {
	log.Printf("Building docker image %s...", name)
	buildCmd := exec.Command("docker", "build", "-t", name, ".")
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	buildCmd.Dir = DOCKERFILE_PATH
	if err := buildCmd.Run(); err != nil {
		log.Fatalf("Failed to build Docker image %s: %v", name, err)
	}
}

func runDockerContainer(name string) {
	log.Printf("Running docker container %s...", name)
	runCmd := exec.Command("docker", "run", "-d", "-p", "8080:8080", name)
	runCmd.Stdout = os.Stdout
	runCmd.Stderr = os.Stderr
	if err := runCmd.Run(); err != nil {
		log.Fatalf("Failed to run Docker container %s: %v", name, err)
	}
}

func copyFileToDockerDirectory(name string) {
	copyCmd := exec.Command("cp", os.Args[1], DOCKERFILE_PATH+"/results.json")
	copyCmd.Stdout = os.Stdout
	copyCmd.Stderr = os.Stderr
	if err := copyCmd.Run(); err != nil {
		log.Fatalf("Failed to copy file to container %s: %v", name, err)
	}
}

func main() {
	var name string = createDockerImageName(os.Args[1])

	copyFileToDockerDirectory(name)
	createDockerImage(name)
	runDockerContainer(name)

	fmt.Println("")
	fmt.Println("Visualiser is running at http://localhost:8080")
	fmt.Println(name)
}

func init() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: ./main <result-file.json>")
	}
	if _, err := os.Stat(os.Args[1]); err != nil {
		log.Fatalf("Error: %v", err)
	}
	if _, err := os.Stat(DOCKERFILE_PATH); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
