package main

import "math"

func reverse(input int) int {
	var reversed int64
	quotient := int64(input)
	remainder := int64(0)
	for {
		quotient, remainder = quotient/10, quotient%10
		reversed = reversed*10 + remainder
		if quotient == 0 {
			break
		}
	}
	if reversed > math.MaxInt32 || reversed < math.MinInt32 {
		return 0
	}
	return int(reversed)
}
