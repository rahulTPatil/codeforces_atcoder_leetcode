package leetcode

import "fmt"

func AddBinary(a string, b string) string {
	aBin := fmt.Sprintf("%08b", a)
	bBin := fmt.Sprintf("%08b", b)

	return aBin + bBin
}
