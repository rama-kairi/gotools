/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/rama-kairi/gotools/internal/utilities"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "'init' command will initialize the project",
	Long: `'init' command will do the following:
		1. Create the .gotools directory in the current directory if not exist
		2. initialize the Git repository with main branch, .gitignore file and README.md file if not exist.
		3. Initialized the go mod file if not exist.
		4. Run go mod tidy if go.mod file is exist.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println("'init' command does not accept any arguments")
			os.Exit(1)
		} else {
			fmt.Println("Initializing the project")
			fmt.Println("Creating .gotools directory in the current directory if not exist")
			utilities.CreateFolder(".gotools")
			fmt.Println("Initializing the Git repository with main branch, .gitignore file and README.md file if not exist.")
			utilities.CheckGitStatus()
			utilities.CreateReadme()
			fmt.Println("Initialized the go mod file if not exist.")
			utilities.InitGoMod()
			fmt.Println("Running go mod tidy if go.mod file is exist.")
			utilities.TidyGoMod()
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
