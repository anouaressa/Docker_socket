package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"os"
)

func main() {
	// Define the Docker socket path
	dockerSocketPath := "/var/run/docker.sock"

	// Check if the Docker socket file exists
	_, err := os.Stat(dockerSocketPath)
	if os.IsNotExist(err) {
		fmt.Printf("Docker socket %s not found. Make sure Docker is running.\n", dockerSocketPath)
		return
	}

	// Create a Docker client to connect to the Docker API via the socket
	cli, err := client.NewClient("unix:///var/run/docker.sock", "", nil, nil)
	if err != nil {
		fmt.Printf("Failed to create a Docker client: %v\n", err)
		return
	}

	// List running containers
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: false})
	if err != nil {
		fmt.Printf("Failed to list containers: %v\n", err)
		return
	}

	if len(containers) > 0 {
		fmt.Println("Running Containers:")
		for _, container := range containers {
			fmt.Printf("Container ID: %s, Name: %s\n", container.ID[:12], container.Names[0])
		}
	} else {
		fmt.Println("No running containers found.")
	}
}
