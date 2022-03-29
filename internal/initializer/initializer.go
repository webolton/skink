package initializer

import (
	"fmt"

	"github.com/webolton/skink/internal/lib"
	"github.com/webolton/skink/internal/models"
	"github.com/webolton/skink/internal/prompts"
)

var configLocation = fmt.Sprintf("%s/.skink.yml", lib.HomeDir()) // default config location

func createConfig() {
	_, defaultConfig, err := prompts.DefaultConfigFile.Run()
	prompts.PromptError(err) // handle prompt error

	newConfig := models.ConfigFileData{ConfigFile: configLocation}

	if defaultConfig == "no" {
		result, err := prompts.CustomConfigFile.Run()
		prompts.PromptError(err) // handle prompt error

		newConfig.ConfigFile = result
	}

	syncedDir, err := prompts.SyncedDir.Run()
	prompts.PromptError(err) // handle prompt error

	parsedDir := lib.ParseDirPath(syncedDir) // parse and validate path to synced directory

	newConfig.SyncedDir = parsedDir

	_, defaultAwsCredsFile, err := prompts.DefaultAwsCredsFile.Run()
	prompts.PromptError(err) // handle prompt error

	if defaultAwsCredsFile == "yes" {
		ami, err := prompts.AMI.Run()
		prompts.PromptError(err) // handle prompt error
		newConfig.AMI = ami

		awsRegion, err := prompts.AwsRegion.Run()
		prompts.PromptError(err) // handle prompt error

		newConfig.AwsRegion = awsRegion

		bucket, err := prompts.Bucket.Run()
		prompts.PromptError(err) // handle prompt error

		newConfig.Bucket = bucket

	} // TODO: store credentials in config file, rather than reading from default
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
