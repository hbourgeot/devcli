/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/hbourgeot/devcli/helpers"
	"github.com/spf13/cobra"
)

// sudoCmd represents the sudo command
var sudoCmd = &cobra.Command{
	Use:   "sudo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		isUpdate, _ := cmd.Flags().GetBool("update")
		if isUpdate {
			return helpers.SudoUpdate()
		}

		isUpgrade, _ := cmd.Flags().GetBool("upgrade")
		if isUpgrade {
			return helpers.SudoUpgrade()
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(sudoCmd)

	sudoCmd.Flags().BoolP("update", "u", false, "Update the packages index")
	sudoCmd.Flags().BoolP("upgrade", "p", false, "Upgrade the packages")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sudoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sudoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
