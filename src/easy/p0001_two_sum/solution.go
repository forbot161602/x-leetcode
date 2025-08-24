package main

// Runtime beats 78.83%, while memory beats 36.20%.
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
