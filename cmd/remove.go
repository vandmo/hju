package cmd

import (
	"fmt"
	"path"

	"github.com/spf13/cobra"
	"github.com/vandmo/hju/core"
	"github.com/vandmo/hju/git"
)

var removeCmd = &cobra.Command{
	Use:   "remove <folder>",
	Short: "Removes a repository by it's folder name",
	Args:  cobra.ExactValidArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		folder := path.Clean(args[0])
		hjuFile, parseErr := core.ParseHjuFile()
		if parseErr != nil {
			return parseErr
		}

		if !hjuFile.ContainsFolder(folder) {
			return fmt.Errorf("Trying to remove %s which isn't managed", folder)
		}

		fmt.Printf("Removing %s\n", folder)
		hjuFile.RemoveByFolder(folder)

		writeErr := hjuFile.Write()
		if writeErr != nil {
			return writeErr
		}

		gitIgnore, gitIgnoreParseErr := git.ParseGitIgnore()
		if gitIgnoreParseErr != nil {
			return gitIgnoreParseErr
		}
		gitIgnore.RemoveRootFolder(folder)
		gitIgnoreWriteErr := gitIgnore.Write()
		if gitIgnoreWriteErr != nil {
			return gitIgnoreWriteErr
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
