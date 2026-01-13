package leetcode

func Palindrome(x int) bool {
	var rev int

	if x < 0 {
		return false
	}

	x2 := x

	for x > 0 {
		rev = (rev * 10) + (x % 10)
		x = x / 10
	}

	return rev == x2
}
