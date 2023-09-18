package helpers

import (
	"strings"
)


//returns true if input has newline in string
func ContainsNewLine(s string) bool {
	if strings.Contains(s, "\\n") {
		return true
	} else {
		return false
	}
}
