package main

import (
	"fmt"
)

func isValid(s string) bool {
	stack := []rune{}

	pair := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if matchingOpen, found := pair[char]; found {
			if len(stack) == 0 || stack[len(stack)-1] != matchingOpen {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}

func main() {
	tests := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
	}

	for _, test := range tests {
		fmt.Printf("isValid(%q) = %v\n", test, isValid(test))
	}
}
