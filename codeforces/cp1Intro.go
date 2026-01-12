package main

import "fmt"

func main() {
	var c1 int
	fmt.Scanln(&c1)
	var c2, c3 int
	fmt.Scanln(&c2, &c3)
	var c4 string
	fmt.Scanln(&c4)

	fmt.Println((c1 + c2 + c3), c4)
}
