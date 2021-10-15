package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vandmo/hju/core"
	"github.com/vandmo/hju/git"
)

var addCmd = &cobra.Command{
	Use:   "add <repository>",
	Short: "Adds a repository",
	Args:  cobra.ExactValidArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		repository := args[0]
		hjuFile, parseErr := core.ParseHjuFileOrNew()
		if parseErr != nil {
			return parseErr
		}

		folderName := git.FolderName(repository)
		if hjuFile.ContainsFolder(folderName) {
			return fmt.Errorf("Trying to add %s which already exists", folderName)
		}

		gitErr := clone(repository)
		if gitErr != nil {
			return gitErr
		}

		hjuFile.Add(repository)

		writeErr := hjuFile.Write()
		if writeErr != nil {
			return writeErr
		}

		gitIgnore, gitIgnoreParseErr := git.ParseGitIgnore()
		if gitIgnoreParseErr != nil {
			return gitIgnoreParseErr
		}
		gitIgnore.AddRootFolderIfNeeded(folderName)
		gitIgnoreWriteErr := gitIgnore.Write()
		if gitIgnoreWriteErr != nil {
			return gitIgnoreWriteErr
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
