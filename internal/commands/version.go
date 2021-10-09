package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of skink",
	Long:  `Prints the version number of skink`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("skink v0.1")
	},
}
