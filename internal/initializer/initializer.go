package initializer

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/webolton/skink/internal/lib"
)

func Execute() {
	newConfigPrompt := promptui.Select{
		Label: "Would you like to create a new skink configuration?",
		Items: []string{"yes", "no"},
	}

	_, result, err := newConfigPrompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "yes" {
		if lib.CheckFile("~/.skink.conf") {
			fmt.Println("there is a file")
		} else {
			fmt.Println("no file")
		}
	}

	if result == "no" {
		fmt.Println("old config")
	}
}
