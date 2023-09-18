package engine

import (
	"strings"
)

type options struct {
	basic      bool
	text       string
	banner     string
	filename   string
	inputError bool
}

func HandleUserInput(args []string) options {
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
