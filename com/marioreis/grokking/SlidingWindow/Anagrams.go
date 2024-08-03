package SlidingWindow

func FindStringAnagrams(str string, pattern string) []int {
	characters := make(map[rune]int)
	windowLeft, matchCount := 0, 0

	for _, c := range pattern {
		characters[rune(c)]++
	}

	resultIndices := make([]int, 0)

	for windowRight := 0; windowRight < len(str); windowRight++ {
		rightChar := rune(str[windowRight])

		_, ok := characters[rightChar]

		if ok {
			characters[rightChar]--

			if characters[rightChar] == 0 {
				matchCount++
			}
		}

		if matchCount == len(characters) {
			resultIndices = append(resultIndices, windowLeft)
		}

		if windowRight >= len(pattern)-1 {
			leftChar := rune(str[windowLeft])
			_, ok := characters[leftChar]

			if ok {
				if characters[leftChar] == 0 {
					matchCount--
				}

				characters[leftChar]++
			}
			windowLeft++
		}
	}

	return resultIndices
}
