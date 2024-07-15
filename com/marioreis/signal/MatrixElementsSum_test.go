package signal

import (
	"strings"
	"testing"
)

func Test_MatrixElementsSum(t *testing.T) {
	matrix := [][]int{{1}, {5}, {0}, {2}}
	println(solution1(matrix))
	s := "teste"
	r := strings.ReplaceAll(s, "t", "")
	println(r)

}

func Test_CountDiffCharacters(t *testing.T) {
	i := findLength("araaci", 1)
	println(i)
}

func findLength(str string, k int) int {
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

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func solution1(matrix [][]int) int {
	sum := 0

	for i := range matrix {
		for j := range matrix[i] {
			if i == 0 {
				sum += matrix[i][j]
			}

			if i >= 1 && matrix[i][j] != 0 && matrix[i-1][j] != 0 {
				sum += matrix[i][j]
			}
		}
	}
	return sum

}
