package leetcode

func Palindrome(x int) bool {
	var ans bool
	var digits []int

	if x >= 0 && x <= 9 {
		ans = true
		return ans
	}

	for x > 0 {
		c := x % 10
		digits = append(digits, c)
		x = x / 10
	}

	for i := 0; i < len(digits)/2; i++ {
		for j := len(digits) - 1; j > len(digits)/2; j-- {
			if digits[i] == digits[j] {
				ans = true
			} else {
				ans = false
			}
		}
	}
	return ans
}
