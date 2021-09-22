package commands

import (
	"github.com/spf13/cobra"
	"github.com/webolton/skink/internal/initializer"
)

var cmdInit = &cobra.Command{
	Use:   "init skink",
	Short: "Setup skink",
	Long: `Provide the directory for skink to watch, the cloud storage and key where
				 skink is to push your files, and a location for this configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		initializer.Execute()
	},
}
