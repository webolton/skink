package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdInit = &cobra.Command{
	Use:   "init skink",
	Short: "Setup skink.",
	Long: `Provide the directory for skink to watch, the cloud storage and key where
				 skink is to push your files, and a location for this configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(" INIT IS CALLED")
	},
}
