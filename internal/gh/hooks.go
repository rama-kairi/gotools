package gh

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/rama-kairi/gotools/internal/utilities"
)

type hookType string

const (
	PreCommit        hookType = "pre-commit"
	PrepareCommitMsg hookType = "prepare-commit-msg"
	PrePush          hookType = "pre-push"
	PostCommit       hookType = "post-commit"
)

// Create file in the root of the project
func createGitHook(hooktype hookType) {
	folder := ".gotools"
	subfolder := ".githooks"
	file := string(hooktype)

	// Check if .gotools folder exists in th root of the project else create it
	utilities.CreateFolder(folder)

	// Check if .gotools folder exists in th root of the project else create it
	utilities.CreateSubFolder(folder, subfolder)

	// Check if "pre-commit" file exists in the .gotools/.githooks/ else create it
	if _, err := os.Stat(folder + "/" + subfolder + "/" + file); os.IsNotExist(err) {
		fmt.Printf("Creating %s file\n", folder+"/"+subfolder+"/"+file)
		file, err := os.Create(folder + "/" + subfolder + "/" + file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		switch hooktype {
		case PreCommit:
			_, err = file.WriteString(preCommitHook)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		case PrepareCommitMsg:
			_, err = file.WriteString(prepareCommitMsgHook)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		case PrePush:
			_, err = file.WriteString(prePushHook)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		case PostCommit:
			_, err = file.WriteString(postCommitHook)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

		default:
			fmt.Println("Unknown hook type")
		}
	}
	fmt.Printf("🪝 %s Hook already available...moving on...(remove the file or the '.githooks' folder and re-run the command for default creation.\n", file)
}

func HookInstaller() {
	goToolGithookDir := ".gotools/.githooks/"
	githookDir := ".git/hooks"

	fmt.Println("")
	fmt.Println("👺 Installing golangci-lint globally...")
	// Install go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	out, err := exec.Command("go", "get", "github.com/golangci/golangci-lint/cmd/golangci-lint@latest").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(string(out))
	}

	fmt.Println("🪝 Installing git hooks...")
	fmt.Println("")
	// Remove all teh files from the githookDir
	if err := os.RemoveAll(githookDir); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create the githookDir
	if err := os.MkdirAll(githookDir, 0o755); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// List all the files in the githookDir
	files, err := ioutil.ReadDir(goToolGithookDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Copy all the files from the githookDir to the githookDir
	for _, file := range files {
		if err := os.Link(goToolGithookDir+file.Name(), githookDir+"/"+file.Name()); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Change the permission of the file
		if err := os.Chmod(githookDir+"/"+file.Name(), 0o755); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("🪝 Installed %s Hook...\n", file.Name())
	}

	fmt.Println("")
	fmt.Println("✅ Git Hook setup done successfully...")
	fmt.Println("")
}

func AllHookInitializer() {
	utilities.CheckGitStatus()
	utilities.CreateSubFolder(".gotools", ".githooks")
	createGitHook(PreCommit)
	createGitHook(PrepareCommitMsg)
	createGitHook(PrePush)
	createGitHook(PostCommit)
	HookInstaller()
}

func HookInitializer(hooktype hookType) {
	utilities.CheckGitStatus()
	utilities.CreateSubFolder(".gotools", ".githooks")
	createGitHook(hooktype)
}
