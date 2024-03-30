package domain

func FindMissingNumberOn(array []int) int {
	n := len(array) + 1
	expectedSum := ((n) * (n + 1)) / 2
	actualSum := 0

	for _, v := range array {
		actualSum += v
	}

	return expectedSum - actualSum
}

func Solution(sequence []int) bool {
	count := 0
	lastNumber := sequence[0] - 1
	m := -1000000

	for i, _ := range sequence {
		if count > 0 && sequence[i] <= m {
			count++
		}

		if sequence[i] <= lastNumber {
			count++
		}

		if sequence[i] > m {
			m = sequence[i]
		}

		lastNumber = sequence[i]

	}

	return count <= 1

}

func FindMissingNumberO2n(array []int) int {
	n := len(array)
	temp := make([]int, n+1)

	for i, _ := range array {
		temp[array[i]-1] = 1
	}

	result := 0

	for i, _ := range temp {
		if temp[i] == 0 {
			result = i + 1
		}
	}

	return result

}
