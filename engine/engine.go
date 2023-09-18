package engine

import (
	"fmt"
	"os"
	"strings"

	"01.kood.tech/git/jsaar/go-reloaded/ascii-art/helpers"
	"01.kood.tech/git/jsaar/go-reloaded/ascii-art/shadow"
	"01.kood.tech/git/jsaar/go-reloaded/ascii-art/standard"
	"01.kood.tech/git/jsaar/go-reloaded/ascii-art/thinkertoy"
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
						banner := standard.Standard(line)
						multilineBanner = append(multilineBanner, banner...)
					}
				}
				outputStr = helpers.CompileBannerString(multilineBanner)
			} else {
				banner := standard.Standard(userInputString)
				outputStr = helpers.CompileBannerString(banner)
			}

			helpers.GenerateFile(outputStr, os.Args[1][9:])
		case "shadow":
			shadow.Shadow(userInputString)
		case "thinkertoy":
			thinkertoy.Thinkertoy(userInputString)
		}
		fmt.Println(userSelectedOption)
	} else {
		fmt.Println("Please enter an argument with or without a font trigger")
		return
	}
}
