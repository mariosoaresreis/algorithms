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

// https://www.youtube.com/watch?v=HbbYPQc-Oo4&ab_channel=Techdose
func FindAllSubArraysWithGivenSum(array []int, k int) int {
	n := len(array)

	if n == 0 {
		return 0
	}

	mpp := map[int]int{}
	currSum := 0
	i := 0
	count := 0

	for i < n {
		currSum += array[i]

		if currSum == k {
			count++
		}

		val, ok := mpp[currSum-k]

		if ok {
			count += val
		}

		mpp[currSum] += 1
		i++
	}

	return count

}

func FindSubArraysPositions(array []int, k int) (start int, end int) {
	start = 0
	end = 0
	sum := 0

	for i := range array {
		sum += array[i]
		end = i

		if sum == k {
			return
		} else {
			for sum > k {
				sum -= array[start]
				start++

				if sum == k {
					return
				}

			}
		}
	}

	return start, end
}
