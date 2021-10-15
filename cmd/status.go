package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vandmo/hju/core"
	"github.com/vandmo/hju/git"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Prints status summary for all managed repositories",
	Args:  cobra.ExactValidArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		hjuFile, parseErr := core.ParseHjuFile()
		if parseErr != nil {
			return parseErr
		}
		for _, folder := range hjuFile.Folders {
			gitErr := doPrintStatus(folder)
			if gitErr != nil {
				return gitErr
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func doPrintStatus(folder string) error {
	ref, err := git.SymbolicRef(folder, "HEAD")
	if err != nil {
		return err
	}
	lastSlashInd := strings.LastIndex(ref, "/")
	branch := strings.TrimSpace(ref[lastSlashInd+1:])
	status, statusErr := git.GetStatus(folder)
	if statusErr != nil {
		return statusErr
	}
	fmt.Printf("\033[32m%s\033[0m (\033[36m%s\033[0m)", folder, branch)
	if status.Tracked > 0 || status.Untracked > 0 {
		fmt.Printf(" \033[31m[T%d,U%d]\033[0m", status.Tracked, status.Untracked)
	}
	fmt.Println()
	return nil
}
