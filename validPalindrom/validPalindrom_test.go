package validPalindrom

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.



Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
*/

func TestValidPalindrom(t *testing.T) {
	testTable := []struct {
		input  string
		output bool
	}{
		{
			input:  "A man, a plan, a canal: Panama",
			output: true,
		},
		{
			input:  "race a car",
			output: false,
		},
		{
			input:  " ",
			output: true,
		},
		{
			input:  "0P",
			output: false,
		},
	}
	for index, test := range testTable {
		t.Run(fmt.Sprintf("%d", index), func(t *testing.T) {
			result := TwoPointerSolution(test.input)
			assert.Equal(t, test.output, result)
		})
	}
}
