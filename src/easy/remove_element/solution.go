package main

func removeElement(numbers []int, target int) int {
	cursor := len(numbers)
	for index, number := range numbers {
		if number != target {
			continue
		}
		for cursor > index {
			cursor--
			if value := numbers[cursor]; value != target {
				numbers[index], numbers[cursor] = value, target
				break
			}
		}
		if cursor == index {
			break
		}
	}
	return cursor
}
