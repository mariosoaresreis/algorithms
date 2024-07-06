package TwoPointers

import (
	"fmt"
)

// Find subarrays with target
func FindSubArrays(arr []int, target int) [][]int {
	result := make([][]int, 0)

	for left := range arr {

		for k := range arr {
			product := 1
			subarray := make([]int, 0)

			for right := left; right <= k; right++ {
				product = product * arr[right]
				subarray = append(subarray, arr[right])
				println(fmt.Sprintf("left=%d", left))
				println(fmt.Sprintf("right=%d.%v", right, arr[right]))
				println(fmt.Sprintf("product=%d", product))
				println(fmt.Sprintf("subarray=%v", subarray))
			}

			if product < target && len(subarray) > 0 {
				result = append(result, subarray)
			}
		}

	}

	return result
}
