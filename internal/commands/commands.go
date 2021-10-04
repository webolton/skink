package commands

func Exeute() {
	rootCmd.AddCommand(cmdInit)
	rootCmd.Execute()
}
