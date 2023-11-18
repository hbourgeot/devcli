package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func SudoUpdate() error {
	distroCmd := exec.Command("uname", "-a")
	output, err := distroCmd.Output()
	if err != nil {
		return err
	}

	distro := string(output)
	fmt.Println(distro)
	if strings.Contains(distro, "Ubuntu") {
		cmd := exec.Command("bash", "-c", "sudo apt update")

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
	}

	return nil
}

func SudoUpgrade() error {
	distroCmd := exec.Command("uname", "-a")
	output, err := distroCmd.Output()
	if err != nil {
		return err
	}

	distro := string(output)
	fmt.Println(distro)
	if strings.Contains(distro, "Ubuntu") {
		cmd := exec.Command("bash", "-c", "sudo apt upgrade")

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
	}

	return nil
}
