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

var configLocation = fmt.Sprintf("%s/.skink.yml", lib.HomeDir()) // default config location

func createConfig() {
	_, defaultConfig, err := prompts.DefaultConfigFile.Run()
	prompts.PromptError(err) // handle prompt error

	newConfig := configFileData{configFile: configLocation}

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

	_, result, err := prompts.NewConfigPrompt.Run() // Prompt new config file?
	prompts.PromptError(err)                        // Handle prompt error

	if result == "yes" {
		if lib.FileExists(configLocation) { // Does a config file already exist?
			_, result, err := prompts.NewConfigConfirmation.Run() // Prompt new config or use old?
			prompts.PromptError(err)                              // handle prompt error

			if result == "no" {
				fmt.Println("Manually check on config for now, but this will be automated")
			}
		} else {
			createConfig()
		}
	}

	if result == "no" {
		if !lib.FileExists(configLocation) {
			fmt.Println("Are you sure? A configuration file does not exist")
		}
		fmt.Println("Manually check on config for now, but this will be automated")
	}
}
