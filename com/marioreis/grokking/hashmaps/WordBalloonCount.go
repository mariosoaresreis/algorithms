package hashmaps

func maxNumberOfBalloons(text string) int {
	wordBalloon := "balloon"
	mapBalloon := make(map[rune]int)
	mapCount := make(map[rune]int)
	smallestCount := 1000000

	for _, v := range wordBalloon {
		mapBalloon[rune(v)]++
	}

	for _, v := range text {
		mapCount[rune(v)]++
	}

	for k, _ := range mapBalloon {
		if mapCount[k] == 0 {
			return 0
		}

		number := mapCount[k] / mapBalloon[k]

		if number < smallestCount {
			smallestCount = number
		}
	}

	return smallestCount
}
