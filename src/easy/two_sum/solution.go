package main

func twoSum(numbers []int, target int) []int {
	compIndexes := map[int]int{}
	for index, number := range numbers {
		if compIndex, ok := compIndexes[number]; ok {
			return []int{compIndex, index}
		}
		compIndexes[target-number] = index
	}
	return []int{}
}
