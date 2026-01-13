package leetcode

import "sort"

func SearchInsert(nums []int, target int) int {
	var pos int
	nums = append(nums, target)

	sort.Ints(nums)

	for i, v := range nums {
		if v == target {
			pos = i
			break
		}
	}
	return pos
}
