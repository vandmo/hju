package cmd

import (
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
			gitErr := git.PrintDivergence(folder, commit)
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
