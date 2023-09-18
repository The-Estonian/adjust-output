package helpers

import (
	"os"
)

// generates a new file with inputString as content and names the file as fileName
func GenerateFile(inputString, filename string) {

	newFile, error := os.Create(filename)
	if error != nil {
		panic("Error creating file")
	} else {
		newFile.WriteString(inputString)
	}
}
