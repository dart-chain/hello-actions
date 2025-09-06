package utils

import (
	"regexp"
)

func IsVersionValid(version string) bool {
	regexStr := `^v(\d+)\.(\d+)\.(\d+)(-[\w\.-]+)?(\+[\w\.-]+)?$`
	matched, _ := regexp.MatchString(regexStr, version)
	return matched
}

func ExtractVersion(version string) string {
	regexStr := `^v(\d+)\.(\d+)\.(\d+)`
	re := regexp.MustCompile(regexStr)
	matches := re.FindStringSubmatch(version)
	if len(matches) > 0 {
		return matches[0]
	}
	return ""
}
