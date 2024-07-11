package SlidingWindow

func FindMaxSumSubArray(k int, arr []int) int {
	maxSum := 0
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]

		if i >= k {
			sum -= arr[i-k]
		}

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}
