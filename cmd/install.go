/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/hbourgeot/devcli/helpers"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "i",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		isGolang, _ := cmd.Flags().GetBool("golang")
		isNpm, _ := cmd.Flags().GetBool("npm")
		isYarn, _ := cmd.Flags().GetBool("yarn")
		isPnpm, _ := cmd.Flags().GetBool("pnpm")
		isDev, _ := cmd.Flags().GetBool("save-dev")
		isGlobal, _ := cmd.Flags().GetBool("global")

		if isGolang {
			if isGlobal {
				return helpers.GoInstall(args...)
			}

			return helpers.Get(args...)
		}

		if isNpm {
			return helpers.NpmInstall(args, isDev)
		}

		if isYarn {
			return helpers.YarnInstall(args, isDev)
		}

		if isPnpm {
			helpers.PnpmInstall(args, isDev)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	installCmd.Flags().BoolP("golang", "g", false, "Get a golang package")
	installCmd.Flags().BoolP("npm", "n", false, "Get a NPM package")
	installCmd.Flags().BoolP("yarn", "y", false, "Get a NPM package with yarn")
	installCmd.Flags().BoolP("pnpm", "p", false, "Get a NPM package with pnpm")
	installCmd.Flags().BoolP("save-dev", "D", false, "Install as dev dependency")
	installCmd.Flags().BoolP("global", "G", false, "install globally (with NPM or 'go install')")
}
