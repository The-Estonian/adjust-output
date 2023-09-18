package engine

import (
	"os"
	"strings"

	"01.kood.tech/git/jsaar/go-reloaded/ascii-art/banners"
	"01.kood.tech/git/jsaar/go-reloaded/ascii-art/helpers"
)

type options struct {
	basic      bool
	text       string
	banner     string
	filename   string
	inputError bool
}

func Start() {
	userInput := handleUserInput(os.Args)
	if userInput.inputError == true {
		helpers.LaunchError()
	}
	if userInput.basic == true {
		// $ go run . "Your Text"
		if helpers.ContainsNewLine(userInput.text) {
			banners.PrintMultilineText(userInput.text, userInput.banner)
		} else {
			banners.PrintText(userInput.text, userInput.banner)
		}
	} else if userInput.basic == false && len(os.Args) == 3 {
		// $ go run . "Your Text" <standard|shadow|thinkertoy>
		if helpers.ContainsNewLine(userInput.text) {
			banners.PrintMultilineText(userInput.text, userInput.banner)
		} else {
			banners.PrintText(userInput.text, userInput.banner)
		}
	} else if userInput.basic == false && len(os.Args) == 4 {
		// $ go run . --output=test.txt "Your Text" <standard|shadow|thinkertoy>
		if helpers.ContainsNewLine(userInput.text) {
			bannerStr := GetMultilineBannerStr(userInput.text, userInput.banner)
			if len(bannerStr) > 0 {
				helpers.GenerateFile(bannerStr, userInput.filename)
			}
		} else {
			bannerArr := banners.EncodeText(userInput.text, userInput.banner)
			outputStr := helpers.CompileBannerString(bannerArr)
			if len(outputStr) > 0 {
				helpers.GenerateFile(outputStr, userInput.filename)
			}
		}
	}
}

func GetMultilineBannerStr(input string, userBanner string) string {
	var multilineBanner []string
	lines := strings.Split(input, "\\n")
	for _, line := range lines {
		banner := banners.EncodeText(line, userBanner)
		multilineBanner = append(multilineBanner, banner...)
	}
	outputStr := helpers.CompileBannerString(multilineBanner)
	return outputStr
}

func handleUserInput(args []string) options {
	var options options
	args = args[1:]
	if len(args) == 1 {
		// $ go run . "Your Text"
		options.inputError = false
		options.basic = true
		options.text = args[0]
		options.banner = "standard"

		if strings.Contains(options.text, "--output=") {
			options.inputError = true
			return options
		}
	} else if len(args) == 2 {
		// $ go run . "Your Text" <standard|shadow|thinkertoy>
		options.inputError = false
		options.basic = false
		options.text = args[0]
		if args[1] == "standard" || args[1] == "shadow" || args[1] == "thinkertoy" {
			options.banner = args[1]
		} else {
			options.inputError = true
			return options
		}
	} else if len(args) == 3 {
		// $ go run . --output=filename.txt "Your Text" <standard|shadow|thinkertoy>
		if !strings.Contains(args[0], "--output=") {
			options.inputError = true
			return options
		}
		options.inputError = false
		options.basic = false
		options.filename = args[0][9:]
		options.text = args[1]
		if args[2] == "standard" || args[2] == "shadow" || args[2] == "thinkertoy" {
			options.banner = args[2]
		} else {
			options.inputError = true
			return options
		}
		options.banner = args[2]
	} else if len(args) > 3 {
		options.inputError = true
		return options
	}
	return options
}
