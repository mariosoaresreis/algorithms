package SlidingWindow

func FindLongestSubstring(str string, k int) int {
	startWindow := 0
	characterMap := make(map[rune]int)
	maxCharactersDiff := k
	maxLength := 0

	for endWindow, rightCharacter := range str {
		characterMap[rightCharacter]++

		if len(characterMap) > 1 && maxCharactersDiff > 0 {
			maxCharactersDiff--
		}

		if maxCharactersDiff == 0 {
			startWindow++
			endWindow = startWindow
			maxCharactersDiff = k
		}

		if (endWindow - startWindow + 1) > maxLength {
			maxLength = endWindow - startWindow + 1
		}
	}

	return maxLength
}
