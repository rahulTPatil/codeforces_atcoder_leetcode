package main

import (
	L "codeforces_atcoder_leetcode/leetcode"
	"fmt"
)

//import A "codeforces_atcoder_leetcode/atcoder"

//import C "codeforces_atcoder_leetcode/codeforces"

func main() {
	digits := []int{1, 3, 5, 6}
	target := 7
	fmt.Println(L.SearchInsert(digits, target))
}
