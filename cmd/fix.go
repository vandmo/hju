package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vandmo/hju/core"
	"github.com/vandmo/hju/git"
)

var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Formats hju.json and fixes .gitignore accordingly",
	Args:  cobra.ExactValidArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		hjuFile, parseErr := core.ParseHjuFile()
		if parseErr != nil {
			return parseErr
		}

		writeErr := hjuFile.Write()
		if writeErr != nil {
			return writeErr
		}

		gitIgnore, gitIgnoreParseErr := git.ParseGitIgnore()
		if gitIgnoreParseErr != nil {
			return gitIgnoreParseErr
		}
		for _, folder := range hjuFile.Folders {
			gitIgnore.AddRootFolderIfNeeded(folder)
		}
		gitIgnoreWriteErr := gitIgnore.Write()
		if gitIgnoreWriteErr != nil {
			return gitIgnoreWriteErr
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(fixCmd)
}
