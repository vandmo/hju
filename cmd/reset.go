package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vandmo/hju/core"
	"github.com/vandmo/hju/git"
)

var resetCmd = &cobra.Command{
	Use:   "reset <commit>",
	Short: "Resets all managed repositories to a specific commit",
	Args:  cobra.ExactValidArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		commit := args[0]
		hjuFile, parseErr := core.ParseHjuFile()
		if parseErr != nil {
			return parseErr
		}
		for _, folder := range hjuFile.Folders {
			gitErr := doReset(folder, commit)
			if gitErr != nil {
				return gitErr
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}

func doReset(folder string, commit string) error {
	fmt.Printf("--- \033[32mResetting to %s in %s\033[0m\n", commit, folder)
	return git.Reset(folder, commit)
}
