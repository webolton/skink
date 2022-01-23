package initializer

import (
	"fmt"

	"github.com/webolton/skink/internal/lib"
	"github.com/webolton/skink/internal/prompts"
)

type configFileData struct {
	configFile string
	syncedDir  string
	bucket     string
	key        string
}

func createConfig() {
	_, defaultConfig, err := prompts.DefaultConfigFile.Run()
	prompts.PromptError(err) // handle prompt error

	newConfig := configFileData{configFile: "~/.skink.yml"}

	if defaultConfig == "no" {
		result, err := prompts.CustomConfigFile.Run()
		prompts.PromptError(err) // handle prompt error

		newConfig.configFile = result
	}

	syncedDir, err := prompts.SyncedDir.Run()
	prompts.PromptError(err) // handle prompt error

	newConfig.syncedDir = syncedDir

	fmt.Println(newConfig)
}

func Execute() {

	_, result, err := prompts.NewConfigPrompt.Run()

	prompts.PromptError(err) // handle prompt error

	if result == "yes" {
		if lib.CheckFile("~/.skink.conf") {
			_, result, err := prompts.NewConfigConfirmation.Run()
			prompts.PromptError(err) // handle prompt error

			if result == "no" {
				fmt.Println("Check existing skink config probably a function")
			}
		} else {
			createConfig()
		}
	}

	if result == "no" {
		fmt.Println("old config")
	}
}
