package signal

func ExecCountMines(matrix [][]bool) [][]int {
	result := make([][]int, 0)

	for j := 0; j < len(matrix); j++ {
		result = append(result, make([]int, len(matrix[j])))
	}

	maxI := len(result) - 1
	maxJ := len(result[0]) - 1

	for i := 0; i <= maxI; i++ {
		for j := 0; j <= maxJ; j++ {
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

	if (j >= 1) && (matrix[i][j-1] == true) {
		count++
	}

	if (j < maxJ) && (matrix[i][j+1] == true) {
		count++
	}

	if i < maxI && (matrix[i+1][j] == true) {
		count++
	}

	if (i < maxI) && (j < maxJ) && (matrix[i+1][j+1] == true) {
		count++
	}

	if (i >= 1) && (j < maxJ) && (matrix[i-1][j+1]) == true {
		count++
	}

	if (i >= 1) && (j >= 1) && (matrix[i-1][j-1] == true) {
		count++
	}

	if (i < maxI) && (j >= 1) && (matrix[i+1][j-1] == true) {
		count++
	}

	return count

}
