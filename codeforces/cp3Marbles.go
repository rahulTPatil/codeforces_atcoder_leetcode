package main

import "fmt"

func main() {
	var num, c1 int
	fmt.Scanf("%d", &c1)
	for c1 > 0 {
		digit := c1 % 10
		if digit == 1 {
			num++
		}
		c1 = c1 / 10
	}
	fmt.Println(num)
}
