package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	var num int
	fmt.Scanln(&num)

	var speed int
	var speed_list []int

	for range num {
		fmt.Scan(&speed)
		speed_list = append(speed_list, speed)
	}

	var speed_list2 = make([]int, len(speed_list))
	copy(speed_list2, speed_list)
	sort.Ints(speed_list2)
	fmt.Printf("%v %v %v\n", (slices.Index(speed_list, speed_list2[0]) + 1), (slices.Index(speed_list, speed_list2[1]) + 1), (slices.Index(speed_list, speed_list2[2]) + 1))
}
