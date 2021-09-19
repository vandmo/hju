package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vandmo/hju/core"
	"github.com/vandmo/hju/git"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clones the needed repositories",
	Args:  cobra.ExactValidArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		hjuFile, parseErr := core.ParseHjuFile()
		if parseErr != nil {
			return parseErr
		}
		for _, repository := range hjuFile.Repositories {
			gitErr := git.Clone(repository)
			if gitErr != nil {
				return gitErr
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
