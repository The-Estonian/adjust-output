package helpers

import (
	"01.kood.tech/git/jsaar/go-reloaded/ascii-art/banners"
	"strings"
)

func GetMultilineBannerStr(input string, userBanner string) string {
	var multilineBanner []string
	lines := strings.Split(input, "\\n")
	for _, line := range lines {
		banner := banners.EncodeText(line, userBanner)
		multilineBanner = append(multilineBanner, banner...)
	}
	outputStr := CompileBannerString(multilineBanner)
	return outputStr
}
