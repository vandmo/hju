package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vandmo/hju/internal"
)

var rootCmd = &cobra.Command{
	Use:     "hju",
	Short:   "Taking git to the next level",
	Version: internal.Version(),
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
