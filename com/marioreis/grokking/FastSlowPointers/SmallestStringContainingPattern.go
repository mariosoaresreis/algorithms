package FastSlowPointers

import "math"

func FindSubstring(str string, pattern string) string {
	//minString := ""
	minSize := math.MaxInt32
	characters := make(map[rune]int)
	subStrStart, leftWindow, match := 0, 0, 0

	for _, char := range pattern {
		characters[rune(char)]++
	}

	for rightWindow, rightChar := range str {
		_, ok := characters[rightChar]

		if !ok {
			characters[rightChar]--

			if characters[rightChar] >= 0 {
				match++
			}
		}

		if match == len(characters) {
			newString := str[leftWindow : rightWindow+1]

			if len(newString) < minSize {
				//minString = newString
				minSize = len(newString)
				subStrStart = leftWindow
			}
		}

		for match == len(characters) {
			subStrStart = leftWindow
			leftChar := str[leftWindow]
			_, ok := characters[rune(leftChar)]

			if ok {
				match--
				characters[rune(leftChar)]++
			}

			leftWindow++
		}

	}

	return str[subStrStart : subStrStart+minSize]
}
