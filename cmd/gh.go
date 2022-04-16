/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/rama-kairi/gotools/internal/gh"
	"github.com/rama-kairi/gotools/internal/utilities"
	"github.com/spf13/cobra"
)

// ghCmd represents the gh command
var ghCmd = &cobra.Command{
	Use:   "gh",
	Short: "GH stands for Git Hooks",
	Long: `GH stands for Git Hooks.
	'gh' is a command line tool that helps you manage your git hooks.
	It is a wrapper around the git hooks that comes with git.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("Too many arguments. You can only specify one of ['pre-commit', 'pre-push', 'auto-setup'] these flags.")
		} else if len(args) == 0 {
			cmd.Help()
		} else {
			arg := strings.ToLower(args[0])

			switch arg {
			case "pre-commit":
				utilities.CreateSubFolder(".gotools", ".githooks")
				gh.HookInitializer(gh.PreCommit)
			case "pre-push":
				gh.HookInitializer(gh.PrePush)
			case "prepare-commit-msg":
				gh.HookInitializer(gh.PrepareCommitMsg)
			case "auto-setup":
				gh.AllHookInitializer()
			case "post-commit":
				gh.HookInitializer(gh.PostCommit)
			case "install":
				gh.HookInstaller()
			default:
				fmt.Println("Error: must specify one of ['pre-commit', 'pre-push', 'auto-setup', 'post-commit','prepare-commit-msg','install'] these flags.")

			}
		}
	},
}

func init() {
	rootCmd.AddCommand(ghCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ghCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ghCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
