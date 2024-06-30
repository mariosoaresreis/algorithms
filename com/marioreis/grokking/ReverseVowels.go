package grokking

import "strings"

func ReverseVowels(str string) string {
	// TODO: Write your code here
	vowels := "AEIOU"
	indexes := make([]int, 0)

	for i := range str {
		c := strings.ToUpper(string(rune(str[i])))

		if strings.Contains(vowels, c) {
			indexes = append(indexes, i)
		}

	}

	offset := 0
	result := []byte(str)
	l := len(indexes) - 1

	for i := 0; i < len(indexes)/2; i++ {
		temp := result[indexes[i]]
		println(l - offset)
		println(i)
		result[indexes[i]] = result[indexes[l-offset]]
		result[indexes[l-offset]] = temp
		offset++
	}

	return string(result)
}
