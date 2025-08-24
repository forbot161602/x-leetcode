package main

// Runtime beats 85.36%, while memory beats 54.97%.
func romanToInt(text string) int {
	chars := []rune(text)
	result := 0
	cursor := 0
	maxCursor := len(chars) - 1
	for {
		value := romanSymbolValues[string(chars[cursor])]
		cursor++
		if cursor <= maxCursor {
			if next := romanSymbolValues[string(chars[cursor])]; value < next {
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
