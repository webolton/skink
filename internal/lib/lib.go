package lib

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func HomeDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return dirname
}

func FileExists(fileName string) bool {
	_, error := os.Stat(fileName)
	return error == nil
}

// Checks provided path for a directory. Returns either a string with an expanded home path, or the
// string provided if it were a legitimate path with an existing directory.
func ParseDirPath(directoryPath string) string {
	tildeMatcher, err := regexp.Compile(`(\~)(.*)`)

	if err != nil {
		log.Fatal(err)
	}

	if tildeMatcher.MatchString(directoryPath) {
		directoryString := tildeMatcher.FindStringSubmatch(directoryPath)[2]
		fullPath := fmt.Sprintf("%s%s", HomeDir(), directoryString)
		if !FileExists(fullPath) {
			err := fmt.Errorf("no directory exists at %s", fullPath)
			log.Fatal(err)
		}
		return fullPath
	}

	if !FileExists(directoryPath) {
		err := fmt.Errorf("no directory exists at %s", directoryPath)
		log.Fatal(err)
	}

	return directoryPath
}
