package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vandmo/hju/core"
	"github.com/vandmo/hju/git"
)

var divergenceCmd = &cobra.Command{
	Use:   "divergence <commit>",
	Short: "Summarizes how the repositories are ahead or behind another commit",
	Args:  cobra.ExactValidArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		commit := args[0]
		hjuFile, parseErr := core.ParseHjuFile()
		if parseErr != nil {
			return parseErr
		}
		for _, folder := range hjuFile.Folders {
			gitErr := doPrintDivergence(folder, commit)
			if gitErr != nil {
				return gitErr
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(divergenceCmd)
}

func doPrintDivergence(folder string, commit string) error {
	hasCommit, hasCommitErr := git.HasCommit(folder, commit)
	if hasCommitErr != nil {
		return hasCommitErr
	}
	if !hasCommit {
		fmt.Printf("%s \033[31m[no-such-commit]\033[0m\n", folder)
		return nil
	}
	divergence, divergenceErr := git.GetDivergence(folder, commit)
	if divergenceErr != nil {
		return divergenceErr
	}
	if divergence.Ahead == 0 && divergence.Behind == 0 {
		fmt.Printf("%s (\033[36mup-to-date\033[0m)", folder)
	} else {
		fmt.Printf("%s \033[33m[A%d,B%d]\033[0m", folder, divergence.Ahead, divergence.Behind)
	}
	fmt.Println()
	return nil
}
