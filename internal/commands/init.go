package commands

import (
	"github.com/spf13/cobra"
	"github.com/webolton/skink/internal/initializer"
)

type promptContent struct {
	errorMsg string
	label    string
}

func init() {
	rootCmd.AddCommand(initializeCmd)
}

var initializeCmd = &cobra.Command{
	Use:   "init",
	Short: "Setup skink",
	Long: `Provide the directory for skink to watch, the cloud storage and key where
skink is to push your files, and a location for this configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		initializer.Execute()
	},
}
