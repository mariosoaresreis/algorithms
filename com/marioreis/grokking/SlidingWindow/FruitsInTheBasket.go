package SlidingWindow

func FindLengthFruitsInTheBasket(arr []string) int {
	maxLength := 0
	frequency := make(map[string]int)
	startWindow := 0

	for endWindow, rightCharacter := range arr {
		frequency[rightCharacter]++

		for len(frequency) > 2 {
			leftCharacter := arr[startWindow]
			frequency[leftCharacter]--

			if frequency[leftCharacter] == 0 {
				delete(frequency, leftCharacter)
			}

			startWindow++
		}

		if (endWindow - startWindow + 1) > maxLength {
			maxLength = endWindow - startWindow + 1
		}
	}

	return maxLength
}
