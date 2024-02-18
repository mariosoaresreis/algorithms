package domain

func consecutiveSubArrays(array []int) [][]int {
	result := make([][]int, 0)

	for start := 0; start < len(array); start++ {
		for groupSize := 1; groupSize <= len(array); groupSize++ {
			subArray := make([]int, 0)
			ok := false

			for j := start; j < groupSize; j++ {
				subArray = append(subArray, array[j])
				ok = true
			}

			if ok {
				result = append(result, subArray)
			}
		}

	}

	return result
}
