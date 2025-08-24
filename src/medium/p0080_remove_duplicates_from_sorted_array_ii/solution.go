package main

// Runtime beats 100.00%, while memory beats 18.63%.
func removeDuplicates(numbers []int) int {
	if len(numbers) == 1 {
		return 1
	}

	index := 0
	count := 1
	value := numbers[0]
	for _, number := range numbers[1:] {
		if number != value {
			index++
			count = 1
			value = number
			numbers[index] = number
			continue
		}
		if count == 1 {
			index++
			count++
			numbers[index] = number
		}
	}
	return index + 1
}
