package domain

func FindMissingNumber(array []int) int {
	n := len(array) + 1
	expectedSum := ((n) * (n + 1)) / 2
	actualSum := 0

	for _, v := range array {
		actualSum += v
	}

	return expectedSum - actualSum
}
