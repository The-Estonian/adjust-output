package engine

import (
	"01.kood.tech/git/jsaar/go-reloaded/ascii-art/banners"
	"01.kood.tech/git/jsaar/go-reloaded/ascii-art/helpers"
	"os"
	"strings"
	"fmt"
)

func Start() {
	if len(os.Args) > 1 {
		userSelectedOption := ""
		userSelectedBanner := "standard"
		userInputString := ""
		fmt.Println(os.Args)
		if len(os.Args) == 2 {
			userInputString = os.Args[1]
			banners.PrintText(userInputString, userSelectedBanner)
		} else if len(os.Args) == 3 && !strings.HasPrefix(os.Args[2], "--output="){
			userInputString = os.Args[1]
			userSelectedBanner := os.Args[2]
			banners.PrintText(userInputString, userSelectedBanner)
		} else {
			userInputString = os.Args[2]
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
			var outputStr string
			var multilineBanner []string
			if helpers.ContainsNewLine(userInputString) {
				lines := strings.Split(userInputString, "\\n")
				for _, line := range lines {
					if line != "\\n" {
						if userSelectedBanner == "shadow" {
							banner := banners.EncodeText(line, userSelectedBanner)
							multilineBanner = append(multilineBanner, banner...)
						} else if userSelectedBanner == "thinkertoy" {
							banner := banners.EncodeText(line, userSelectedBanner)
							multilineBanner = append(multilineBanner, banner...)
						} else {
							banner := banners.EncodeText(line, userSelectedBanner)
							multilineBanner = append(multilineBanner, banner...)
						}
					}
				}
				outputStr = helpers.CompileBannerString(multilineBanner)
			} else {
				banner := banners.EncodeText(userInputString, userSelectedBanner)
				outputStr = helpers.CompileBannerString(banner)
			}
			if len(outputStr) > 0 {
				helpers.GenerateFile(outputStr, userSelectedOption)
			}
		}
	} else {
		helpers.LaunchError()
		return
	}
}
