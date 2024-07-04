package TwoPointers

import "sort"

func SearchTriplets(arr []int, target int) int {
	count := 0
	sort.Ints(arr)

	for i := 0; i <= len(arr)-3; i++ {
		left := i + 1
		right := len(arr) - 1

		for left < right {
			sumAll := arr[i] + arr[left] + arr[right]

			if sumAll < target {
				count++
				newLeft := left + 1

				for newLeft < right {
					nSumAll := arr[i] + arr[newLeft] + arr[right]

					if nSumAll < target {
						count++
					}
					newLeft++
				}
			}

			right--
		}

	}
	return count
}
