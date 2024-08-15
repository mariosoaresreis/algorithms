package hashmaps

func LongestPalindrome(s string) int {
	mapWordCount := make(map[rune]int)
	biggestSumOdd := 0
	sumEven := 0

	for _, v := range s {
		mapWordCount[rune(v)]++
	}

	for k, _ := range mapWordCount {
		if mapWordCount[k]%2 == 0 {
			sumEven += mapWordCount[k]
		} else {
			if mapWordCount[k] > biggestSumOdd {
				biggestSumOdd = mapWordCount[k]
			} else {
				sumEven += mapWordCount[k] - 1
			}
		}
	}

	return sumEven + biggestSumOdd
}
