package helpers

import (
	"os"
	"os/exec"
)

func Get(args ...string) error {
	cmd := exec.Command("go", "get -u")
	cmd.Args = append(cmd.Args, args...)

	// Set command stdout and stderr to os stdout and stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start the command
	err := cmd.Start()
	if err != nil {
		return err
	}

	// Wait for the command to finish
	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

func GoInstall(args ...string) error {
	cmd := exec.Command("go", "install")
	cmd.Args = append(cmd.Args, args...)

	// Set command stdout and stderr to os stdout and stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start the command
	err := cmd.Start()
	if err != nil {
		return err
	}

	// Wait for the command to finish
	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}
