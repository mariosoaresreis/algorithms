package domain

import "fmt"

func canPaint(matrix [][]int, color int, row int, col int, maxRow int, maxCol int, visited [][]int, originalColor int) {
	fmt.Println(row)
	fmt.Println(col)
	matrix[row][col] = color
	visited[row][col] = 1

	if col+1 <= maxCol && visited[row][col+1] != 1 && matrix[row][col+1] == originalColor {
		matrix[row][col+1] = color
		visited[row][col+1] = 1
		canPaint(matrix, color, row, col+1, maxRow, maxCol, visited, originalColor)
	}

	if col-1 >= 0 && visited[row][col-1] != 1 && matrix[row][col-1] == originalColor {
		matrix[row][col-1] = color
		visited[row][col-1] = 1
		canPaint(matrix, color, row, col-1, maxRow, maxCol, visited, originalColor)
	}

	if row+1 <= maxRow && visited[row+1][col] != 1 && matrix[row+1][col] == originalColor {
		matrix[row+1][col] = color
		visited[row+1][col] = 1
		canPaint(matrix, color, row+1, col, maxRow, maxCol, visited, originalColor)
	}

	if row-1 >= 0 && visited[row-1][col] != 1 && matrix[row-1][col] == originalColor {
		matrix[row-1][col] = color
		visited[row-1][col] = 1
		canPaint(matrix, color, row-1, col, maxRow, maxCol, visited, originalColor)
	}

	if col-1 >= 0 && row-1 >= 0 && visited[row-1][col-1] != 1 && matrix[row-1][col-1] == originalColor {
		matrix[row-1][col-1] = color
		visited[row-1][col-1] = 1
		canPaint(matrix, color, row-1, col-1, maxRow, maxCol, visited, originalColor)
	}

	if col+1 <= maxCol && row+1 <= maxRow && visited[row+1][col+1] != 1 && matrix[row+1][col+1] == originalColor {
		matrix[row+1][col+1] = color
		visited[row+1][col+1] = 1
		canPaint(matrix, color, row+1, col+1, maxRow, maxCol, visited, originalColor)
	}

	if row-1 >= 0 && col+1 <= maxCol && visited[row-1][col+1] != 1 && matrix[row-1][col+1] == originalColor {
		matrix[row-1][col+1] = color
		visited[row-1][col+1] = 1
		canPaint(matrix, color, row-1, col+1, maxRow, maxCol, visited, originalColor)
	}

	if row+1 <= maxRow && col-1 >= 0 && visited[row+1][col-1] != 1 && matrix[row+1][col-1] == originalColor {
		matrix[row+1][col-1] = color
		visited[row+1][col-1] = 1
		canPaint(matrix, color, row+1, col-1, maxRow, maxCol, visited, originalColor)
	}

}

func TestPaint() {
	var matrix [][]int = make([][]int, 10)
	var visited [][]int = make([][]int, 10)

	matrix[0] = []int{0, 2, 3, 4, 5, 6, 7, 8, 9, 1}
	matrix[1] = []int{1, 0, 0, 0, 5, 6, 7, 8, 9, 1}
	matrix[2] = []int{1, 2, 3, 0, 5, 6, 7, 8, 9, 1}
	matrix[3] = []int{1, 2, 3, 0, 5, 0, 7, 8, 9, 1}
	matrix[4] = []int{1, 2, 3, 0, 0, 6, 7, 8, 9, 1}
	matrix[5] = []int{1, 2, 3, 4, 0, 6, 7, 8, 9, 1}
	matrix[6] = []int{1, 2, 3, 4, 0, 0, 0, 0, 0, 0}
	matrix[7] = []int{1, 2, 3, 0, 5, 6, 7, 8, 9, 1}
	matrix[8] = []int{1, 2, 3, 0, 5, 6, 7, 8, 9, 1}
	matrix[9] = []int{0, 0, 0, 0, 5, 6, 7, 8, 9, 1}

	visited[0] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	visited[1] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	visited[2] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	visited[3] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	visited[4] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	visited[5] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	visited[6] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	visited[7] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	visited[8] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	visited[9] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	canPaint(matrix, 12, 0, 0, 9, 9, visited, matrix[0][0])
	fmt.Println(matrix)

}
