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
