package prompts

import "github.com/manifoldco/promptui"

var NewConfigPrompt = promptui.Select{
	Label: "Would you like to create a new skink configuration?",
	Items: []string{"yes", "no"},
}
