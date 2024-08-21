package SlidingWindow

func DistinctCharacters(word string, k int) int {
	characters := make(map[rune]int)
	left := 0
	match := 0
	countWords := 0

	for right := 0; right < len(word); right++ {
		rightChar := rune(word[right])
		characters[rightChar]++

		if characters[rightChar] == 1 {
			match++
		}

		if match == k {
			countWords++
		}

		if right-left+1 >= k {
			leftChar := rune(word[left])

			if characters[leftChar] == 1 {
				match--
			}

			characters[leftChar]--
			left++
		}
	}

	return countWords
}
