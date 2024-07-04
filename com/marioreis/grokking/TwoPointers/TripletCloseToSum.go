package TwoPointers

import (
	"fmt"
	"math"
	"sort"
)

func SearchTriplet(arr []int, targetSum int) int {
	sort.Ints(arr)
	maxIndex := len(arr) - 1
	smallestDiff := math.MaxInt32

	for i := 0; i <= len(arr)-3; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		left := i + 1
		right := maxIndex

		for left < right {
			println(fmt.Sprintf("i=%v;%v", i, arr[i]))
			println(fmt.Sprintf("left=%v;%v", left, arr[left]))
			println(fmt.Sprintf("right=%v;%v", right, arr[right]))
			sumAll := arr[i] + arr[left] + arr[right]
			remaining := targetSum - sumAll

			if remaining == 0 {
				return targetSum
			}

			if (abs(remaining) < abs(smallestDiff)) ||
				(abs(remaining) == abs(smallestDiff) && remaining > smallestDiff) {
				smallestDiff = remaining
			}

			if remaining < 0 {
				right--
			}

			if remaining > 0 {
				left++
			}
		}
	}

	return targetSum - smallestDiff
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}

	return n
}
