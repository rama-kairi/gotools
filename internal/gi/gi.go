package gi

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rama-kairi/gotools/internal/utilities"
)

type LangType string

const (
	GoLang LangType = "go"
)

const gitIgnoreLinkPrefix = "https://raw.githubusercontent.com/github/gitignore/main/"

func GenerateGitIgnore(lang LangType) {
	fmt.Println("Generating .gitignore file for " + lang + " projects")

	// check if .git exists
	utilities.CheckGitStatus()

	// Create .gitignore file in the current directory
	file, err := os.Create(".gitignore")
	if err != nil {
		panic(err)
	}

	// Write to file
	switch lang {
	case GoLang:
		_, err := file.WriteString(`# .gitignore\n
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
