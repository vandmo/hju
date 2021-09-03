package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vandmo/hju/core"
)

var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Formats hju.json and fixes .gitignore accordingly",
	RunE: func(cmd *cobra.Command, args []string) error {
		repositories, parseErr := core.ParseHjuFile()
		if parseErr != nil {
			return parseErr
		}

		writeErr := core.WriteHjuFile(repositories)
		if writeErr != nil {
			return writeErr
		}

		gitIgnore, gitIgnoreParseErr := core.ParseGitIgnore()
		if gitIgnoreParseErr != nil {
			return gitIgnoreParseErr
		}
		for _, repo := range repositories {
			gitIgnore.AddIfNeeded(repo)
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
