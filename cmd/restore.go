package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vandmo/hju/core"
	"github.com/vandmo/hju/git"
)

var restoreCmd = &cobra.Command{
	Use:   "restore <pathspec>",
	Short: "Restores all managed repositories",
	Args:  cobra.ExactValidArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pathspec := args[0]
		hjuFile, parseErr := core.ParseHjuFile()
		if parseErr != nil {
			return parseErr
		}
		for _, folder := range hjuFile.Folders {
			gitErr := doRestore(folder, pathspec)
			if gitErr != nil {
				return gitErr
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)
}

func doRestore(folder string, pathspec string) error {
	fmt.Printf("--- \033[32mRestoring %s in %s\033[0m\n", pathspec, folder)
	return git.Restore(folder, pathspec)
}
