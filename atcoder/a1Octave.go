package main

import (
	"fmt"
	"math"
)

func main() {
	var ci1, ci2 int
	fmt.Scanln(&ci1, &ci2)
	fmt.Println((ci1 * int(math.Pow(2, float64(ci2)))))
}
