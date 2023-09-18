package engine

import (
	"os"

	"01.kood.tech/git/jsaar/go-reloaded/ascii-art/banners"
	"01.kood.tech/git/jsaar/go-reloaded/ascii-art/helpers"
)


// starts the ENGINE!, handles input, parses it and prints it to the console or file
func Start() {
	userInput := HandleUserInput(os.Args)
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
			bannerStr := helpers.GetMultilineBannerStr(userInput.text, userInput.banner)
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
