package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)
	if num > 2 && num%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
