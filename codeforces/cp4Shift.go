package main

import (
	"fmt"
)

func divideByTwo(arr []int) []int {
	for i := range arr {
		arr[i] = arr[i] / 2
	}
	return arr
}

func main() {
	var count, digit, operations int
	fmt.Scan(&count)
	var digits []int

	for range count {
		fmt.Scan(&digit)
		digits = append(digits, digit)
	}
Outerloop:
	for {
		digits = divideByTwo(digits)
		operations++
		for i := range digits {
			if digits[i]%2 != 0 {
				break Outerloop
			}
		}
	}
	fmt.Println(operations)
}
