package main

// https://leetcode.com/problems/roman-to-integer/
func romanToInt(text string) int {
	chars := []rune(text)
	result := 0
	cursor := 0
	maxCursor := len(chars) - 1
	for {
		value := RomanSymbolValues[string(chars[cursor])]
		cursor++
		if cursor <= maxCursor {
			if next := RomanSymbolValues[string(chars[cursor])]; value < next {
				value = next - value
				cursor++
			}
		}
		result += value
		if cursor > maxCursor {
			break
		}
	}
	return result
}

var RomanSymbolValues = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}
