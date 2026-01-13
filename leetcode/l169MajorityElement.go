package leetcode

func MajorityElement(nums []int) int {
	res := 0
	maj := 0

	for _, n := range nums {
		if maj == 0 {
			res = n
		}
		if n == res {
			maj++
		} else {
			maj--
		}
	}
	return res
}
