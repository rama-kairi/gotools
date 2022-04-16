package utilities

import (
	"fmt"
	"os"
)

// Create folder in the root of the project
func CreateFolder(folder string) {
	// Check if .gotools folder exists in th root of the project else create it
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		fmt.Printf("Creating %s folder\n", folder)
		if err := os.Mkdir(folder, 0o755); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Printf("📁 %s folder already exists...moving on...\n", folder)
}

// Create Subfolder in the root of the project
func CreateSubFolder(folder string, subfolder string) {
	// Check if .gotools folder exists in th root of the project else create it
	CreateFolder(folder)

	// Check if .gotools folder exists in th root of the project else create it
	if _, err := os.Stat(folder + "/" + subfolder); os.IsNotExist(err) {
		fmt.Printf("Creating %s folder\n", folder+"/"+subfolder)
		if err := os.Mkdir(folder+"/"+subfolder, 0o755); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Printf("📁 %s folder already exists...moving on...\n", subfolder)
}

// Create file in the root of the project
func CreateFile(folder string, subfolder string, file string) {
	// Check if .gotools folder exists in th root of the project else create it
	CreateFolder(folder)

	// Check if .gotools folder exists in th root of the project else create it
	CreateSubFolder(folder, subfolder)

	// Check if "pre-commit" file exists in the .gotools/.githooks/ else create it
	if _, err := os.Stat(folder + "/" + subfolder + "/" + file); os.IsNotExist(err) {
		fmt.Printf("Creating %s file\n", folder+"/"+subfolder+"/"+file)
		file, err := os.Create(folder + "/" + subfolder + "/" + file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()
	}
}
