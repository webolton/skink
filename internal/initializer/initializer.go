package initializer

import (
	"fmt"

	"github.com/webolton/skink/internal/lib"
	"github.com/webolton/skink/internal/prompts"
)

func Execute() {

	_, result, err := prompts.NewConfigPrompt.Run()

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
