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

	for range count {
		if scanner.Scan() {
			entries = append(entries, scanner.Text())
		}
	}

	for _, line := range entries {
		l := len(line)
		if l > 10 {
			fmt.Printf("%s%v%s\n", string(line[0]), (l - 2), string(line[l-1]))
		} else {
			fmt.Println(line)
		}
	}

}
