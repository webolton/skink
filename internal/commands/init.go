package commands

import (
	"github.com/spf13/cobra"
	"github.com/webolton/skink/internal/initializer"
)

func init() {
	rootCmd.AddCommand(initializeCmd)
}

var initializeCmd = &cobra.Command{
	Use:   "init",
	Short: "Setup skink",
	Long:  `Setup and configure skink.`,
	Run: func(cmd *cobra.Command, args []string) {
		initializer.Execute()
	},
}
