package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vandmo/hju/core"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clones the needed repositories",
	RunE: func(cmd *cobra.Command, args []string) error {
		repositories, parseErr := core.ParseHjuFile()
		if parseErr != nil {
			return parseErr
		}
		for _, repo := range repositories {
			fmt.Println("Cloning: " + repo)
			gitErr := core.GitIt("clone", repo)
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
