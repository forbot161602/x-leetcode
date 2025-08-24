package main

// Runtime beats 100%, while memory beats 45.95%.
func mySqrt(number int) int {
	if number <= 1 {
		return number
	}
	lower, upper := 1, number/2
	for lower <= upper {
		middle := lower + (upper-lower)/2
		if middle*middle <= number {
			lower = middle + 1
		} else {
			upper = middle - 1
		}
	}
	return upper
}
