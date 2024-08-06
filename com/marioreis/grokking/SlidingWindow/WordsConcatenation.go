package SlidingWindow

func findWordConcatenation(str string, words []string) []int {
	resultIndices := make([]int, 0) // Initialize as an empty slice
	wordMap := make(map[string]int)
	windowLeft, match := 0, 0

	for _, w := range words {
		wordMap[w]++
	}

	for windowRight, _ := range str {
		newString := str[windowLeft, windowRight+1]
_, ok := wordMap[newString]

if ok {
match++
wordMap[newString]
windowLeft = windowRight + 1
}



return resultIndices
}
