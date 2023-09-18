package engine

import (
	"fmt"
	"os"
	"strings"
	"01.kood.tech/git/jsaar/go-reloaded/ascii-art/helpers"
	"01.kood.tech/git/jsaar/go-reloaded/ascii-art/banners"
	
)

func Start() {
	if len(os.Args) > 1 {
		userSelectedOption := ""
		userSelectedBanner := "standard"
		userInputString := ""
		if len(os.Args) == 2 {
			userInputString = os.Args[1]
		} else {
			userInputString = os.Args[2]
		}
		for i := 0; i < len(os.Args); i++ {
			if strings.HasPrefix(os.Args[i], "--output=") {
				userSelectedOption = os.Args[1][9:]
			} else if strings.HasPrefix(os.Args[i], "standard") {
				userSelectedBanner = "standard"
			} else if strings.HasPrefix(os.Args[i], "shadow") {
				userSelectedBanner = "shadow"
			} else if strings.HasPrefix(os.Args[i], "thinkertoy") {
				userSelectedBanner = "thinkertoy"
			}
		}

		switch userSelectedBanner {
		case "standard":
			var outputStr string
			var multilineBanner []string
			if helpers.ContainsNewLine(userInputString) {
				lines := strings.Split(userInputString, "\\n")
				for _, line := range lines {
					if line != "\\n" {
						banner := banners.EncodeText(line, "standard")
						multilineBanner = append(multilineBanner, banner...)
					}
				}
				outputStr = helpers.CompileBannerString(multilineBanner)
			} else {
				banner := banners.EncodeText(userInputString, "standard")
				outputStr = helpers.CompileBannerString(banner)
			}

			helpers.GenerateFile(outputStr, os.Args[1][9:])
		case "shadow":
			banners.EncodeText(userInputString, "shadow")
		case "thinkertoy":
			banners.EncodeText(userInputString, "thinkertoy")
		}
		fmt.Println(userSelectedOption)
	} else {
		fmt.Println("Please enter an argument with or without a font trigger")
		return
	}
}
