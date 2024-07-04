package TwoPointers

import (
	"fmt"
	"sort"
	"testing"
)

func SearchTriplets(arr []int) [][]int {
	sort.Ints(arr)
	triplets := make([][]int, 0)

	for i := 0; i < len(arr); i++ {
		fixed := i
		right := len(arr) - 1
		left := fixed + 1

		for left < right {
			if arr[fixed] == (arr[left]+arr[right])*-1 {
				triplets = append(triplets, []int{arr[fixed], arr[left], arr[right]})
				left++
				continue
			}

			if (arr[fixed])*-1 > (arr[left] + arr[right]) {
				left++
			}

			if (arr[fixed])*-1 < arr[left]+arr[right] {
				right--
			}
		}
	}

	// TODO: Write your code here
	return triplets
}

func Test_SearchTriplet(t *testing.T) {
	result := SearchTriplets([]int{-3, 0, 1, 2, -1, 1, -2})
	print(fmt.Sprintf("%v", result))
}
