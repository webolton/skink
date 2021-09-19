package commands

import (
	"github.com/spf13/cobra"

)

func Execute() {
	var rootCmd = &cobra.Command{Use: "skink"}
	rootCmd.AddCommand(cmdInit)
	rootCmd.Execute()
}
