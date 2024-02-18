package domain

import "fmt"

func ArrayCombinations(slice []int) {
	result := make([][]int, 0)

	for i := 0; i < len(slice); i++ {
		combine(0, slice, []int{})
	}

	fmt.Println(result)
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

		if i+1 < maxSize {
			combine(i+1, originalSlice, tempSlice)
		}
	}

	return finalResult
}
