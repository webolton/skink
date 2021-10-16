package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
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

		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		version := os.Getenv("VERSION")
		fmt.Println("skink " + version)
	},
}
