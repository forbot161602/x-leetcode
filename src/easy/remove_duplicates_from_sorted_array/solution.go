package main

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
