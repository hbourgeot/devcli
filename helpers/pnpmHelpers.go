package helpers

import (
	"os"
	"os/exec"
)

func PnpmInstall(packages []string, isDev, isGlobal bool) error {
	// Create npm command with the i argument for install
	var cmd *exec.Cmd
	if len(packages) == 0 {
		cmd = exec.Command("pnpm", "i")
	} else {
		cmd = exec.Command("pnpm", "add")
		cmd.Args = append(cmd.Args, packages...)
	}

	// Check if install is for dev packages
	if isDev {
		cmd.Args = append(cmd.Args, "-D")
	} else if isGlobal {
		return NpmInstall(packages, false, true)
	}

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

	// Return success
	return nil
}

func PnpmInit() error {
	// Initialize npm package
	cmd := exec.Command("npm", "init", "-y")

	// Redirect stdout and stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start command execution
	err := cmd.Start()
	if err != nil {
		return err
	}

	// Wait for command to finish
	err = cmd.Wait()
	if err != nil {
		return err
	}

	// Return success
	return nil
}

func PnpmUninstall(packages []string) error {
	cmd := exec.Command("npm", "uninstall")

	if len(packages) > 0 {
		cmd.Args = append(cmd.Args, packages...)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// start the command
	err := cmd.Start()
	if err != nil {
		return err
	}

	// wait for the command to finish
	err = cmd.Wait()
	if err != nil {
		return err
	}

	// return success
	return nil
}
