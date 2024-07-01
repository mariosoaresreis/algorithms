package grokking

func MoveElements(arr []int) int {
	count := 1
	lastNumber := arr[0]

	if len(arr) == 1 {
		return 1
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] != lastNumber {
			count++
			lastNumber = arr[i]
		}
	}

	return count
}
