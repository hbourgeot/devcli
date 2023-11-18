/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hbourgeot/devcli/helpers"
	"github.com/spf13/cobra"
)

// pCmd represents the p command
var pCmd = &cobra.Command{
	Use:   "p",
	Short: "PNPM Command",
	Long: `Runs a PNPM command, like:
	[a]dd
	[i]nstall (default)
	[u]ninstall
	i[n]it

	With the --save-dev or -D flag and packages for install/uninstall`,
	RunE: func(cmd *cobra.Command, args []string) error {
		action, _ := cmd.Flags().GetString("action")
		isDev, _ := cmd.Flags().GetBool("save-dev")

		fmt.Println(action)
		switch action {
		case "a", "add":
			if err := helpers.PnpmInstall(args, isDev); err != nil {
				return err
			}
		case "i", "install":
			if err := helpers.PnpmInstall(nil, false); err != nil {
				return err
			}
		case "u", "uninstall":
			if err := helpers.PnpmUninstall(args); err != nil {
				return err
			}
		case "n", "init":
			if err := helpers.PnpmInit(); err != nil {
				return err
			}
		default:
			if action == "init" {
				if err := helpers.PnpmInit(); err != nil {
					return err
				}
			}
			if err := helpers.PnpmInstall(args, isDev); err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	pCmd.Flags().StringP("action", "a", "a", "PNPM Action, can be: [a]dd, i[n]it, [i]nstall, [u]ninstall")
	pCmd.Flags().BoolP("save-dev", "D", false, "Is a Dev dependency?")

	rootCmd.AddCommand(pCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
