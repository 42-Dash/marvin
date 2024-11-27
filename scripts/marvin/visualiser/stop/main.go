package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const DOCKER_CONTAINER_NAME = "visualiser"

func stopDockerContainer(name string) {
	log.Printf("Stopping docker container %s...", name)
	stopCmd := exec.Command("docker", "stop", name)
	stopCmd.Stdout = os.Stdout
	stopCmd.Stderr = os.Stderr
	if err := stopCmd.Run(); err != nil {
		log.Fatalf("Failed to stop Docker container %s: %v", name, err)
	}
}

func removeTheContainer(name string) {
	log.Printf("Removing docker container %s...", name)
	removeCmd := exec.Command("docker", "rm", name)
	removeCmd.Stdout = os.Stdout
	removeCmd.Stderr = os.Stderr
	if err := removeCmd.Run(); err != nil {
		log.Fatalf("Failed to remove Docker container %s: %v", name, err)
	}
}

func main() {
	stopDockerContainer(DOCKER_CONTAINER_NAME)
	removeTheContainer(DOCKER_CONTAINER_NAME)

	fmt.Println("\n\033[32mVisualiser stopped and removed.\033[0m")
}
