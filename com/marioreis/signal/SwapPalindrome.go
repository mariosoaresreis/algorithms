package signal

import "strings"

/*
Given a string, find out if its characters can be rearranged to form a palindrome.

Example

For inputString = "aabb", the output should be
solution(inputString) = true.

We can rearrange "aabb" to make "abba", which is a palindrome.

Input/Output

[execution time limit] 4 seconds (go)

[memory limit] 1 GB

[input] string inputString

A string consisting of lowercase English letters.

Guaranteed constraints:
1 ≤ inputString.length ≤ 50.

[output] boolean

true if the characters of the inputString can be rearranged to form a palindrome, false otherwise.*/

func SwapPalindrome(inputString string) bool {
	allowedOdds := 0
	odds := 0

	if len(inputString)%2 != 0 {
		allowedOdds = 1
	}

	for len(inputString) > 0 {
		sizeBefore := len(inputString)
		c := inputString[0]
		inputString = strings.ReplaceAll(inputString, string(c), "")
		sizeAfter := len(inputString)
		result := sizeAfter - sizeBefore

		if result%2 != 0 {
			odds++
		}

		if odds > allowedOdds {
			return false
		}

	}

	return true

}
