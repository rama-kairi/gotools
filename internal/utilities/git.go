package utilities

import (
	"fmt"
	"os"
	"os/exec"
)

func scanAnswers() string {
	fmt.Println("Do you want GoTools to initialize git for you? (y/n)")
	answer := ""
	_, err := fmt.Scan(&answer)
	if err != nil {
		fmt.Println(err)
	}

	// if args are more than 1 or the answer is not 'y' or 'n', ask again
	if answer != "y" && answer != "n" {
		fmt.Println("Error: must specify one of ['y', 'n']")
		fmt.Println("Info: Or just press 'ctrl+c' to exit.")
		scanAnswers()
	}

	return answer
}

func CheckGitStatus() {
	_, err := os.Stat(".git")
	if os.IsNotExist(err) {
		fmt.Println("git is not initialized, please run 'git init -b main' first.")

		// Do you want GoTools to initialize git for you? (y/n)
		// if yes, run 'git init -b main'
		// if no, exit
		switch scanAnswers() {
		case "y":
			fmt.Println("Running 'git init -b main'")
			out, err := exec.Command("git", "init", "-b", "main").Output()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				fmt.Println(string(out))
			}
		case "n":
			fmt.Println("Exiting...")
			fmt.Println("Please run 'git init -b main' first, then run this command again.")
			os.Exit(1)

			fmt.Println("Running contains")
		}
	}
}
