package signal

import (
	"testing"
)

func Test_ExecCountMines(t *testing.T) {
	matrix := [][]bool{
		{false, false, false},
		{false, false, false}}
	ExecCountMines(matrix)

	//println(len(matrix))
	//println(len(matrix[0]))
	/*
		matrix := [][]int{
			{0, 1, 2},
			{3, 4, 5},
		}

		result := make([][]int, 0)
		for j := 0; j < len(matrix); j++ {
			result = append(result, make([]int, len(matrix[j])))
		}

		for i := 0; i < len(result); i++ {
			for j := 0; j < len(result[0]); j++ {
				println(fmt.Sprintf("%d,%d", i, j))
				println(matrix[i][j])
			}
		}
	*/
	//println(fmt.Sprintf("%v", matrix))
	//println(fmt.Sprintf("%v", result))*/

}

func ExecCountMines(matrix [][]bool) [][]int {
	result := make([][]int, 0)

	for j := 0; j < len(matrix); j++ {
		result = append(result, make([]int, len(matrix[j])))
	}

	maxI := len(result) - 1
	maxJ := len(result[0]) - 1
	for i := 0; i < maxI; i++ {
		for j := 0; j < maxJ; j++ {
			result[i][j] = countMines(matrix, i, j, maxI, maxJ)

		}
	}
	return result
}

func countMines(matrix [][]bool, i int, j int, maxI int, maxJ int) int {
	count := 0

	if (i >= 1) && (matrix[i-1][j] == true) {
		count++
	}
	println("A")

	if (j >= 1) && (matrix[i][j-1] == true) {
		count++
	}
	println("B")

	if (j < maxJ) && (matrix[i][j+1] == true) {
		count++
	}

	println("C")

	if (i < len(matrix[0])-1) && (matrix[i+1][j] == true) {
		count++
	}
	println("D")

	if (i < maxI) && (j < maxJ) && (matrix[i+1][j+1] == true) {
		count++
	}
	println("E")

	if (i >= 1) && (j < maxJ) && (matrix[i-1][j+1]) == true {
		count++
	}
	println("F")

	if (i >= 1) && (j >= 1) && (matrix[i-1][j-1] == true) {
		count++
	}
	println("G")

	if (i < maxI) && (j >= 1) && (matrix[i+1][j-1] == true) {
		count++
	}
	println("H")

	return count

}
