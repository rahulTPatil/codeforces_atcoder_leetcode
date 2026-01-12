package leetcode

import "strings"

func IndexOfString(haystack string, needle string) int {
	if strings.Contains(haystack, needle) {
		return strings.Index(haystack, needle)
	}
	return -1
}
