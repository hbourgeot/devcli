/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// dkerbCmd represents the dkerb command
var dkerbCmd = &cobra.Command{
	Use:   "dkerb",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		command := exec.Command("docker", "build")
		command.Args = append(command.Args, args...)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		// Start the command
		err := command.Start()
		if err != nil {
			return err
		}

		// Wait for the command to finish
		err = command.Wait()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(dkerbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dkerbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only RunE when this command
	// is called directly, e.g.:
	// dkerbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
