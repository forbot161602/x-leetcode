package main

// https://leetcode.com/problems/two-sum/description/
func twoSum(numbers []int, target int) []int {
	compIndexMap := map[int]int{}
	for index, number := range numbers {
		if compIndex, ok := compIndexMap[number]; ok {
			return []int{compIndex, index}
		}
		compIndexMap[target-number] = index
	}
	return []int{}
}
