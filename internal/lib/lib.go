package lib

import "os"

func CheckFile(fileName string) bool {
	_, error := os.Stat(fileName)
	return error == nil
}
