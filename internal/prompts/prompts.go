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

var DefaultAwsCredsFile = promptui.Select{
	Label: "Would you like to use your default AWS credentials file?",
	Items: []string{"yes", "no"},
}

var AMI = promptui.Prompt{
	Label: "The profile of the credentials to use (default or AMI)",
}

var AWSRegion = promptui.Select{
	Label: "aws region",
	Items: []string{"us-east-2", "us-east-1", "us-west-1", "us-west-2"},
}

var Bucket = promptui.Prompt{
	Label: "aws bucket name",
}

// TODO: don't use default AWS credentials file
// var AwsAccessKeyID = promptui.Prompt{
// 	Label: "aws access key id",

// var AwsSecretAccessKey = promptui.Prompt{
// 	Label: "aws secret access key",
// }
