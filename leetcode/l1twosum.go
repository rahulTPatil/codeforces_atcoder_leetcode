package leetcode

func TwoSum(nums []int, target int) []int {
	var sums []int
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]+nums[i] == target {
				sums = append(sums, i, j)
			}
		}
	}
	return sums
}
