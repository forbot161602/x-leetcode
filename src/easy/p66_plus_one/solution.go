package main

// Runtime beats 100%, while memory beats 35.34%.
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] == 10 {
			digits[i] = 0
			continue
		}
		return digits
	}
	return append([]int{1}, digits...)
}
