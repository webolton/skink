package initializer

import (
	"fmt"

	"github.com/webolton/skink/internal/lib"
	"github.com/webolton/skink/internal/prompts"
)

func Execute() {

	_, result, err := prompts.NewConfigPrompt.Run()

	prompts.PromptError(err) // handle prompt error

	if result == "yes" {
		if lib.CheckFile("~/.skink.conf") {
			_, result, err := prompts.NewConfigConfirmation.Run()
			prompts.PromptError(err) // handle prompt error

			if result == "yes" {
				fmt.Println("Call the new config file function")
			}

			if result == "no" {
				fmt.Println("Check existing skink config")
			}
		}
	}

	if result == "no" {
		fmt.Println("old config")
	}
}
