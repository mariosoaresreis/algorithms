package grokking

import "strings"

func IsAnagram(s, t string) bool {
	// TODO: Write your code here

	if len(s) != len(t) {
		return false
	}

	lT := len(t)
	lS := len(s)

	for len(s) > 0 {
		c := string(rune(s[0]))

		if strings.Contains(t, c) {
			s = strings.ReplaceAll(s, c, "")
			t = strings.ReplaceAll(t, c, "")

			if lT-len(s) != lS-len(t) {
				return false
			}

			lT = len(t)
			lS = len(s)

		} else {
			return false
		}
	}

	return len(t) == 0 && len(s) == 0

}
