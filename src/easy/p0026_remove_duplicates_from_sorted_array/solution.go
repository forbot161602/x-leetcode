package main

// Runtime beats 95.11%, while memory beats 7.98%.
func removeDuplicates(numbers []int) int {
	cursor := 0
	values := map[int]bool{}
	for _, number := range numbers {
		if _, ok := values[number]; ok {
			continue
		}
		values[number] = true
		numbers[cursor] = number
		cursor++
	}
	return cursor
}
