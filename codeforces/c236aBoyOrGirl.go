package main

import "fmt"

func main() {
	var s string
	fmt.Scanln(&s)
	var count int
	for i := 0; i < len(s); i++ {
		for j := 1; j < len(s)-1; j++ {
			if s[i] == s[j] {
				count++
			}
		}
	}
	fmt.Println(count)
}
