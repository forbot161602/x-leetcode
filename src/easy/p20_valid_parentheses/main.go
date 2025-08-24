package main

import "fmt"

// https://leetcode.com/problems/valid-parentheses/
func main() {
	solution := isValid
	fmt.Println(solution("()[]{}"))
}

type ParenthesesInfo struct {
	Type    string
	IsStart bool
}

var parenthesesInfos = map[string]ParenthesesInfo{
	"(": {"()", true},
	")": {"()", false},
	"[": {"[]", true},
	"]": {"[]", false},
	"{": {"{}", true},
	"}": {"{}", false},
}
