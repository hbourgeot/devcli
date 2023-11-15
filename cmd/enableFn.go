/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// enableFnCmd represents the enableFn command
var enableFnCmd = &cobra.Command{
	Use:   "enable-fn",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("enableFn called")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(enableFnCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// enableFnCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only RunE when this command
	// is called directly, e.g.:
	// enableFnCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
