package tdd

import (
	"strings"
	"testing"
)

func Test_MatrixElementsSum(t *testing.T) {
	matrix := [][]int{{1}, {5}, {0}, {2}}
	println(solution(matrix))
	s := "teste"
	r := strings.ReplaceAll(s, "t", "")
	println(r)

}

func solution(matrix [][]int) int {
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
