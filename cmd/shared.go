package cmd

import (
	"fmt"
	"os"

	"github.com/vandmo/hju/git"
)

func clone(url string) error {
	folderName := git.FolderName(url)
	if _, err := os.Stat(folderName); err == nil {
		fmt.Println("--- \033[33mNOT cloning: " + url + "\033[0m")
		return nil
	}
	fmt.Println("--- \033[32mCloning: " + url + "\033[0m")
	return git.Clone(url)
}
