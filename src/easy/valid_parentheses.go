package main

// Tag: Stack
// https://leetcode.com/problems/valid-parentheses/
func isValid(text string) bool {
	stack := make([]string, len(text), len(text))
	cursor := -1
	for _, char := range text {
		sign := string(char)
		info := ParenthesesInfos[sign]
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

var ParenthesesInfos = map[string]ParenthesesInfo{
	"(": {"()", true},
	")": {"()", false},
	"[": {"[]", true},
	"]": {"[]", false},
	"{": {"{}", true},
	"}": {"{}", false},
}

type ParenthesesInfo struct {
	Type    string
	IsStart bool
}
