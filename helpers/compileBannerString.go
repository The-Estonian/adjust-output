package helpers

func CompileBannerString(banner []string) string {
	compiledString := ""
	for i, line := range banner {
		if i != 0 || i != (len(banner)+1) {
			compiledString += "\n"
		}
		compiledString += line
	}
	return compiledString
}
