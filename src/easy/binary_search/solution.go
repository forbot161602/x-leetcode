package main

func search(numbers []int, target int) int {
	lower, upper := 0, len(numbers)-1
	for lower <= upper {
		middle := lower + (upper-lower)/2
		number := numbers[middle]
		if number == target {
			return middle
		} else if number < target {
			lower = middle + 1
		} else {
			upper = middle - 1
		}
	}
	return -1
}
