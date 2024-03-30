package domain

func MaxWater(arr []int, n int) int {
	// To store the maximum water
	// that can be stored
	res := 0

	// For every element of the array
	// except first and last element
	for i := 1; i < n-1; i++ {
		// Find maximum element on its left
		left := arr[i]
		for j := 0; j < i; j++ {
			left = max(left, arr[j])
		}

		// Find maximum element on its right
		right := arr[i]
		for j := i + 1; j < n; j++ {
			right = max(right, arr[j])
		}

		// Update maximum water value
		res += max(left, right) - arr[i]
	}
	return res
}

func MaxWater2(arr []int, start int, res int) int {
	if start >= len(arr)-1 {
		return res
	}

	result := 0
	left := 0
	right := 0
	mostRightIndex := 0
	found := false

	for i := start; i < len(arr); i++ {
		if arr[i] > 0 {
			left = arr[i]

			for j := i + 1; j < len(arr); j++ {
				if arr[j] > right {
					right = arr[j]
					mostRightIndex = j
					found = true
				}

				if found {
					break
				}
			}

			right = min(left, right)
			result += (mostRightIndex - start - 1) * right

			if found {
				break
			}
		}
	}

	return MaxWater2(arr, mostRightIndex, result+res)
}
