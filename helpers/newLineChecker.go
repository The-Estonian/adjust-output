package helpers

import (
	"strings"
)

func ContainsNewLine(s string) bool {
	if strings.Contains(s, "\\n") {
		return true
	} else {
		return false
	}
}
