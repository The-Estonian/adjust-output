package helpers

import (
	"os"
)

func GenerateFile(inputString, filename string) {
	
	newFile, error := os.Create(filename)
	if error != nil {
		panic("Error creating file")
	} else {
		newFile.WriteString(inputString)
	}
}
