package main

// Runtime beats 100%, while memory beats 16.88%.
func isValid(text string) bool {
	stack := make([]string, len(text))
	cursor := -1
	for _, char := range text {
		sign := string(char)
		info := parenthesesInfos[sign]
		if info.IsStart {
			cursor += 1
			stack[cursor] = info.Type
			continue
		}
		if cursor < 0 || stack[cursor] != info.Type {
			return false
		}
		stack[cursor] = ""
		cursor -= 1
	}
	return cursor == -1
}
