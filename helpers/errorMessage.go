package helpers

import (
	"fmt"
	"os"
)

func LaunchError() {
	fmt.Println(`
Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard`)
	os.Exit(0)
}
