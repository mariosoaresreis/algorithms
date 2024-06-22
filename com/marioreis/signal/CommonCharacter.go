package signal

import "strings"

func solution(s1 string, s2 string) int {
	word1 := s1
	word2 := s2
	common := 0

	for len(word1) > 0 {

		lW1 := len(word1)
		lW2 := len(word2)
		s := string(word1[0])
		word1 = strings.ReplaceAll(word1, s, "")
		word2 = strings.ReplaceAll(word2, s, "")

		diff1 := lW1 - len(word1)
		diff2 := lW2 - len(word2)

		if diff1 < diff2 {
			common += diff1
		} else {
			common += diff2
		}

	}

	return common
}
