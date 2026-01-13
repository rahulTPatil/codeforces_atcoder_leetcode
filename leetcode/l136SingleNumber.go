package leetcode

func SingleNumber(nums []int) int {
	var ans int

	if len(nums) == 1 {
		return nums[0]
	}

	counts := make(map[int]int)

	for _, item := range nums {
		counts[item]++
	}

	for i, v := range counts {
		if v == 1 {
			ans = i
		}
	}
	return ans
}
