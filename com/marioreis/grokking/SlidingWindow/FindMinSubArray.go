package SlidingWindow

import "math"

func FindMinSubArray(S int, arr []int) int {
	minLength := math.MaxInt32
	sum := 0
	windowStart := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]

		for sum >= S {
			length := i - windowStart + 1

			if length < minLength {
				minLength = length
			}

			sum -= arr[windowStart]
			windowStart++
		}
	}

	if minLength == math.MaxInt32 {
		return 0
	}

	return minLength
}
