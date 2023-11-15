/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hbourgeot/devcli/helpers"
	"github.com/spf13/cobra"
)

// nCmd represents the n command
var nCmd = &cobra.Command{
	Use:   "n",
	Short: "NPM Command",
	Long: `Runs a NPM command, like:
	[i]nstall
	[u]ninstall
	i[n]it

	With the --save-dev or -D flag and packages for install/uninstall
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		action, _ := cmd.Flags().GetString("action")
		isDev, _ := cmd.Flags().GetBool("save-dev")

		fmt.Println(action)
		switch action {
		case "i", "install":
			if err := helpers.NpmInstall(args, isDev); err != nil {
				return err
			}
		case "u", "uninstall":
			if err := helpers.NpmUninstall(args); err != nil {
				return err
			}
		case "n", "init":
			if err := helpers.NpmInit(); err != nil {
				return err
			}
		default:
			if action == "init" {
				if err := helpers.NpmInit(); err != nil {
					return err
				}
			}
			if err := helpers.NpmInstall(args, isDev); err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	nCmd.Flags().StringP("action", "act", "i", "NPM Action, can be: i[n]it, [i]nstall, [u]ninstall")
	nCmd.Flags().BoolP("save-dev", "D", false, "Is a Dev dependency?")

	rootCmd.AddCommand(nCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
