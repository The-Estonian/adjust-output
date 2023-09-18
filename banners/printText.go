package banners

import (
	"fmt"
)

func PrintText(input string, userBanner string) {
	output := EncodeText(input, userBanner)
	for i := 0; i < len(output); i++ {
		fmt.Println(output[i])
	}
}
