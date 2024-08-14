package hashmaps

func largestUniqueNumber(A []int) int {
	mapA := make(map[int]int)
	maxValue := -1

	for _, v := range A {
		mapA[v]++
	}

	for k, v := range mapA {
		if v == 1 && k > maxValue {
			maxValue = k
		}
	}

	return maxValue
}
