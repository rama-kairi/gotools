package utilities

import (
	"fmt"
	"os"
	"os/exec"
)

// Initialized the go mod file if not exist.
func InitGoMod() {
	// check if go.mod file exist
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		// Check if user would like to use go modules and github to store the project?
		// If yes, then create go.mod file
		if ans := scannerWithConfirm("Would you like to use go modules and github to store the project?"); ans == "y" {
			fmt.Println("Creating go.mod file")

			// Ask github username
			username := scanner("Please enter your github username:")

			// Ask for the Project Name
			projectName := scanner("Please enter the project name:")

			// Create go.mod file
			out, err := exec.Command("go", "mod", "init", "github.com/"+username+"/"+projectName).Output()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				fmt.Println(string(out))
			}
		} else {
			projectName := scanner("Please enter the project name:")
			out, err := exec.Command("go", "mod", "init", projectName).Output()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				fmt.Println(string(out))
			}
		}
	}
}

// Run go mod tidy if go.mod file is exist.

func TidyGoMod() {
	// check if go.mod file exist
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		fmt.Println("Running go mod tidy")
		out, err := exec.Command("go", "mod", "tidy").Output()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println(string(out))
		}
	} else {
		fmt.Println("go.mod file not exist")
	}
}
