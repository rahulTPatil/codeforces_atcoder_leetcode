package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var count int
	fmt.Scanln(&count)
	scanner := bufio.NewScanner(os.Stdin)
	var entries []string

	for i := 0; i < count; i++ {
		if scanner.Scan() {
			entries = append(entries, scanner.Text())
		}
	}
}
