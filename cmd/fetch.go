package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vandmo/hju/core"
	"github.com/vandmo/hju/git"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetches all managed repositories",
	Args:  cobra.ExactValidArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		hjuFile, parseErr := core.ParseHjuFile()
		if parseErr != nil {
			return parseErr
		}
		for _, folder := range hjuFile.Folders {
			gitErr := git.Fetch(folder)
			if gitErr != nil {
				return gitErr
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
