package TwoPointers

import (
	"fmt"
	"sort"
)

func MinimumWindowSort(arr []int) int {
	orderedArr := make([]int, len(arr))
	copy(orderedArr, arr)
	sort.Ints(orderedArr)
	start := -1
	end := -1
	maximum := len(arr) - 1
	println(fmt.Sprintf("arr=%v", arr))
	println(fmt.Sprintf("ordered=%v", orderedArr))
	diff := 0

	for i := 0; i < len(arr); i++ {
		if start < 0 && orderedArr[i] != arr[i] {
			start = i
			diff++
		}

		if end < 0 && orderedArr[maximum] != arr[maximum] {
			end = maximum
			diff++
		}

		maximum--
	}

	if diff == 0 {
		return 0
	}

	if start < 0 {
		start = 0
	}

	if end < 0 {
		end = len(arr) - 1
	}

	println(fmt.Sprintf("start=%v", start))
	println(fmt.Sprintf("end=%v", end))
	println(fmt.Sprintf("res=%v", arr[start:end]))

	return len(arr[start : end+1])
}
