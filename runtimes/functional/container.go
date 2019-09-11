package functional

import (
	"errors"
	"os/exec"
	"time"

	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared/schema"
)

func ContainerBootstrap(imageName string) error {
	cmd := exec.Command("singularity", "instance.start", "docker://Nerv-Core-Developers/test1", "test1")
	if err := cmd.Start(); err != nil {
		return err
	}
	return nil
}

// Exec create a container to run function
func ContainerExec(funcID string, data []string) (*schema.RuntimeOutput, time.Duration, error) {
	start := time.Now()
	err := ContainerBootstrap(funcID)
	if err != nil {
		return nil, time.Since(start), errors.New("Failed to bootstrap function " + funcID)
	}

	fData := append([]string{"exec", "-c", "-e", "-p", "instance://test1", "/run.sh"}, data...)
	cmd := exec.Command("singularity", fData...)

	stdout, err := cmd.Output()
	if err != nil {
		return nil, time.Since(start), err
	}
	_ = cmd.Wait()
	fOutput, err := readOutput(stdout)
	if err != nil {
		return nil, time.Since(start), err
	}
	return fOutput, time.Since(start), nil
}
