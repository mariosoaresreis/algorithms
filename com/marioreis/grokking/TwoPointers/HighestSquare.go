package TwoPointers

func MakeSquares(arr []int) []int {
	n := len(arr)
	squares := make([]int, n)
	i := 0
	j := len(arr) - 1
	highestIdx := len(arr) - 1

	for i <= j {
		leftSquare := arr[i] * arr[i]
		rightSquare := arr[j] * arr[j]

		if leftSquare >= rightSquare {
			squares[highestIdx] = leftSquare
			i++
		}

		if leftSquare < rightSquare {
			squares[highestIdx] = rightSquare
			j--
		}
		highestIdx--
	}

	return squares
}
