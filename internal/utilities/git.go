package utilities

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func scanner(question string) string {
	fmt.Println(question, "(y/n)")
	answer := ""
	_, err := fmt.Scan(&answer)
	if err != nil {
		fmt.Println(err)
	}

	// if args are more than 1 or the answer is not 'y' or 'n', ask again
	if answer != "y" && answer != "n" {
		fmt.Println("Error: must specify one of ['y', 'n']")
		fmt.Println("Info: Or just press 'ctrl+c' to exit.")
		scanner(question)
	}

	return answer
}

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

// Generate .gitingore file from -> const gitIgnoreLinkPrefix = "https://raw.githubusercontent.com/github/gitignore/main/" if not exist

func GenerateGitIgnore() {
	const gitIgnoreLinkPrefix = "https://raw.githubusercontent.com/github/gitignore/main/"

	// Create .gitignore file in the current directory if not exist
	if _, err := os.Stat(".gitignore"); os.IsNotExist(err) {
		file, err := os.Create(".gitignore")
		if err != nil {
			panic(err)
		}

		// Write to file
		_, err = file.WriteString(`# .gitignore\n
	# Created by gotools\n`)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		// Read the content of this https://raw.githubusercontent.com/github/gitignore/main/Go.gitignore url
		// and write it to the file
		resp, err := http.Get(gitIgnoreLinkPrefix + "Go.gitignore")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		_, err = file.Write(data)
		if err != nil {
			panic(err)
		}
	}
}

// Crete README.md file if not exist
func CreateReadme() {
	if _, err := os.Stat("README.md"); os.IsNotExist(err) {
		file, err := os.Create("README.md")
		if err != nil {
			panic(err)
		}

		// Write to file
		_, err = file.WriteString(`# Projects README\n
	# Created by gotools\n`)
		if err != nil {
			panic(err)
		}

		defer file.Close()
	}
}
