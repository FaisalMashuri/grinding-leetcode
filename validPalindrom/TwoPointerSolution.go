package validPalindrom

import (
	"fmt"
	"strings"
	"unicode"
)

func TwoPointerSolution(input string) bool {
	input = strings.ToLower(cleanString(input))
	fmt.Println(input)
	if input == "" {
		return true
	}
	left := 0
	right := len(input) - 1
	for left < right {
		if input[left] != input[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func cleanString(s string) string {
	var cleaned string
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			cleaned += string(char)
		}
	}
	return cleaned
}
