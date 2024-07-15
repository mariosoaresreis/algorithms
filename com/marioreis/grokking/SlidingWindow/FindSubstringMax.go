package SlidingWindow

func FindLength(str string, k int) int {
	frequency := make(map[rune]int)
	maxLength := 0
	startWindow := 0

	for windowEnd, rightChar := range str {
		frequency[rightChar]++

		for len(frequency) > k {
			startCharacter := rune(str[startWindow])
			frequency[startCharacter]--

			if frequency[startCharacter] == 0 {
				delete(frequency, startCharacter)
			}

			startWindow++

		}

		maxLength = max(windowEnd+1-startWindow, maxLength)
	}

	return maxLength
}
