package main

import "fmt"

func main() {
	var c1, c2 int
	fmt.Scanln(&c1, &c2)
	if c1*c2%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
