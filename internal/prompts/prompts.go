package prompts

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func PromptError(err error) {
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
}

var NewConfigPrompt = promptui.Select{
	Label: "Would you like to create a new skink configuration?",
	Items: []string{"yes", "no"},
}

var NewConfigConfirmation = promptui.Select{
	Label: "You already have a skink configuration. Would you like to overwrite it and start over?",
	Items: []string{"yes", "no"},
}

var DefaultConfigFile = promptui.Select{
	Label: "Would you like to use the default config file location: ~/.skink.yml (recommended)?",
	Items: []string{"yes", "no"},
}

var CustomConfigFile = promptui.Prompt{
	Label: "Provide a file path and name your skink config",
}

var SyncedDir = promptui.Prompt{
	Label: "The file path to the directory you would like to sync", // TODO: verify existence
}

var Bucket = promptui.Prompt{
	Label: "The s3 bucket to sync files to",
}

var Key = promptui.Prompt{
	Label: "The s3 bucket key for authenticating with AWS", // TODO: verify ability to write
}
