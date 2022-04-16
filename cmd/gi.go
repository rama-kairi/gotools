/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/rama-kairi/gotools/internal/gi"
	"github.com/spf13/cobra"
)

// giCmd represents the gi command
var giCmd = &cobra.Command{
	Use:   "gi",
	Short: "Git Ignore - Generate .gitignore file",
	Long: `Create a .gitignore file based on the provided arguments
	use 'gotools gi go' to create a .gitignore file for golang projects
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			if err := cmd.Help(); err != nil {
				fmt.Println(err)
			}
		} else {
			if args[0] == "go" {
				fmt.Println("Generating .gitignore file for golang projects")

				gi.GenerateGitIgnore(gi.GoLang)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(giCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// giCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// giCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
