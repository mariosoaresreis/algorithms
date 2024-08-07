package SlidingWindow

// FindPermutation find if there is a permutation of the pattern in the str
func FindPermutation(str string, pattern string) bool {
	characters := make(map[rune]int)
	leftWindow, match := 0, 0

	for _, r := range pattern {
		characters[r]++
	}

	for rightWindow, rightChar := range str {
		_, ok := characters[rightChar]

		if ok {
			characters[rightChar]--

			if characters[rightChar] == 0 {
				match++
			}
		}

		if match == len(characters) {
			return true
		}

		if rightWindow >= len(pattern)-1 {
			leftChar := rune(str[leftWindow])
			_, ok := characters[leftChar]
			leftWindow++

			if ok {
				if characters[leftChar] == 0 {
					match--
				}

				characters[leftChar]++
			}

		}

	}

	return false
}
