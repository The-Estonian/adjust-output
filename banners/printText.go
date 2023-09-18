package banners

import (
	"fmt"
	"strings"
)

func PrintText(input string, userBanner string) {
	output := EncodeText(input, userBanner)
	for i := 0; i < len(output); i++ {
		fmt.Println(output[i])
	}
}

func PrintMultilineText(input string, userBanner string) {
	lines := strings.Split(input, "\\n")
	for _, line := range lines {
		PrintText(line, userBanner)
		fmt.Println()

	}

}
