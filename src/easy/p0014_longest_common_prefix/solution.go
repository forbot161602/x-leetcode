package main

import "math"

// Runtime beats 100%, while memory beats 13.31%.
func longestCommonPrefix(texts []string) string {
	prefix := []rune(texts[0])
	for _, text := range texts[1:] {
		chars := []rune(text)
		prefix = prefix[:int(math.Min(float64(len(prefix)), float64(len(chars))))]
		for i := 0; i < len(prefix); i++ {
			if prefix[i] != chars[i] {
				prefix = prefix[:i]
				break
			}
		}
		if len(prefix) == 0 {
			break
		}
	}
	return string(prefix)
}
