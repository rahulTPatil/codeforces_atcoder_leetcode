package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var num int
	x := 0
	fmt.Scanln(&num)
	scanner := bufio.NewScanner(os.Stdin)
	var operations []string

	for range num {
		if scanner.Scan() {
			operations = append(operations, scanner.Text())
		}
	}

	for i := range operations {
		if strings.Contains(operations[i], "++") {
			x++
		} else if strings.Contains(operations[i], "--") {
			x--
		}
	}
	fmt.Println(x)

}
