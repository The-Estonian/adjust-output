package helpers

import (
	"fmt"
	"os"
)

//prints out error message in console
func LaunchError() {
	fmt.Println(`
Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard`)
	os.Exit(0)
}
