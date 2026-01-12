package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s1 string
	fmt.Scanln(&s1)
	runes := []rune(s1)

	if len(runes) > 0 {
		runes[0] = unicode.ToUpper(runes[0])
	}

	s1 = string(runes)
	fmt.Println(s1)
}
