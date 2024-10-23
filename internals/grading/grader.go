package grading

import (
	"archive/tar"
	"bytes"
	"context"
	"dashinette/pkg/parser"
	"fmt"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

// copyToContainer copies the files from the srcPath to the destPath in the container.
func copyToContainer(ctx context.Context, cli *client.Client, containerID, srcPath, destPath string) error {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)

	err := filepath.Walk(srcPath, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.Mode().IsRegular() {
			data, err := os.ReadFile(file)
			if err != nil {
				return err
			}

			header := &tar.Header{
				Name:    filepath.ToSlash(file),
				Mode:    int64(fi.Mode().Perm()),
				Size:    fi.Size(),
				ModTime: fi.ModTime(),
			}
			if err := tw.WriteHeader(header); err != nil {
				return err
			}
			if _, err := tw.Write(data); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	if err := tw.Close(); err != nil {
		return err
	}

	return cli.CopyToContainer(ctx, containerID, destPath, buf, container.CopyToContainerOptions{})
}

func runDockerContainer(team parser.Team, repo string, tracesfile string) error {
	ctx := context.Background()

	client, err := client.NewClientWithOpts(
		client.FromEnv, client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return err
	}

	dir, _ := os.Getwd()
	config := parser.SerializeTesterConfig(team, repo, tracesfile)
	containerConfig := &container.Config{
		Image:      os.Getenv("DOCKER_IMAGE_NAME"),
		Cmd:        []string{"sh", "-c", fmt.Sprintf("./tester '%v'", config)},
		WorkingDir: "/app",
	}
	hostConfig := &container.HostConfig{
		Binds:      []string{fmt.Sprintf("%s/traces:/app/traces", dir)},
		AutoRemove: false,
	}

	resp, err := client.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "")
	if err != nil {
		return err
	}

	err = copyToContainer(ctx, client, resp.ID, repo, "/app")
	if err != nil {
		return err
	}

	// start the container
	if err := client.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return err
	}

	// wait for the container to finish
	statusCh, errCh := client.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-statusCh:
	}

	output, err := client.ContainerLogs(ctx, resp.ID, container.LogsOptions{ShowStdout: false, ShowStderr: true})
	if err != nil {
		return err
	}

	inspect, err := client.ContainerInspect(ctx, resp.ID)
	if err != nil {
		return err
	}
	if inspect.State.ExitCode != 0 {
		return err
	}

	// if err := client.ContainerRemove(ctx, resp.ID, container.RemoveOptions{}); err != nil {
	// 	return err
	// }

	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	_, err = stdcopy.StdCopy(stdout, stderr, output)
	if err != nil {
		return err
	}

	if stderr.Len() > 0 {
		return fmt.Errorf("stderr: %s", stderr.String())
	}

	return nil
}

// grades the assignment for the given team.
// the function returns an error if an error occurred, otherwise nil.
// the function creates a log file with the results of the grading.
func ContainerizedGrader(team parser.Team, repo string, filename string) error {
	// delete file if it exists
	if _, err := os.Stat(filename); err == nil {
		if err := os.Remove(filename); err != nil {
			return fmt.Errorf("failed to delete file: %v", err)
		}
	}

	err := runDockerContainer(team, repo, filename)
	if _, err := os.Stat(filename); err == os.ErrNotExist {
		return fmt.Errorf("cannot create log file")
	}

	return err
}
