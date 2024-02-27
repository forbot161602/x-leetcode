package main

import "fmt"

// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
// Runtime beats 100.00%, while memory beats 18.63%.
func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
}
