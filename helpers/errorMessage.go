package helpers

import (
	"fmt"
)

func LaunchError() {
	fmt.Println(`
Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard`)
}
