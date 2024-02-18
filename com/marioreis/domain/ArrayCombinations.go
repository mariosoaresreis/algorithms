package domain

import "fmt"

func ArrayCombinations(slice []int) {
	combine(0, slice, []int{})
}

func combine(start int, originalSlice []int, subSlice []int) [][]int {
	maxSize := len(originalSlice)

	if start > maxSize-1 {
		return nil
	}

	if maxSize <= len(subSlice) {
		return nil
	}

	finalResult := make([][]int, 0)

	for i := start; i < maxSize; i++ {
		tempSlice := make([]int, len(subSlice))
		copy(tempSlice, subSlice)
		tempSlice = append(tempSlice, originalSlice[i])
		finalResult = append(finalResult, tempSlice)
		fmt.Println(tempSlice)

		if i+1 < maxSize {
			combine(i+1, originalSlice, tempSlice)
		}
	}

	return finalResult
}
