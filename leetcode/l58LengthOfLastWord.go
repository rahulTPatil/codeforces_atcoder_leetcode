package leetcode

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	s2 := strings.Split(s, " ")
	return len(s2[len(s2)-1])
}
