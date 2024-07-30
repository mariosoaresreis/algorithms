package SlidingWindow

func FindLength2(arr []int, k int) int {
	maxLength, windowStart, maxOnesCount := 0, 0, 0

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		if arr[windowEnd] == 1 {
			maxOnesCount++
		}

		if windowEnd-windowStart+1-maxOnesCount > k {
			if arr[windowStart] == 1 {
				maxOnesCount--
			}
			windowStart++
		}

		maxLength = max(maxLength, windowEnd-windowStart+1)
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
