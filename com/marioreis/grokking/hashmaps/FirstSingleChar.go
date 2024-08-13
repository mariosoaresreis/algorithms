package hashmaps

func firstUniqChar(sStr string) int {
	mapString := make(map[rune]int)

	for _, c := range sStr {
		mapString[c]++
	}

	for i, c := range sStr {
		if mapString[c] == 1 {
			return i
		}
	}

	return -1
}
