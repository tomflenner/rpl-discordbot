package utils

import "strings"

const (
	BadUrl     int = -1
	CustomUrl  int = 0
	Steam64Url int = 1

	CustomUrlPrefix  string = "https://steamcommunity.com/id/"
	Steam64UrlPrefix string = "https://steamcommunity.com/profiles/"
)

func GetSteam64FromUrl(url string) string {
	return strings.ReplaceAll(strings.ReplaceAll(url, Steam64UrlPrefix, ""), "/", "")
}

func GetCustomIdFromUrl(url string) string {
	return strings.ReplaceAll(strings.ReplaceAll(url, CustomUrlPrefix, ""), "/", "")
}

func GetUrlType(url string) int {
	if strings.HasPrefix(url, CustomUrlPrefix) {
		return CustomUrl
	} else if strings.HasPrefix(url, Steam64UrlPrefix) {
		return Steam64Url
	} else {
		return BadUrl
	}
}
