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
