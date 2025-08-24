package main

import "fmt"

// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
func main() {
	solution := removeDuplicates
	fmt.Println(solution([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
}
