package domain

import "fmt"

func ArrayCombinations(slice []int) {
	finalResult := make([][]int, 0)
	combine(0, slice, []int{}, &finalResult)
	fmt.Println(finalResult)
}

func combine(start int, originalSlice []int, subSlice []int, finalResult *[][]int) {
	maxSize := len(originalSlice)

	if start > maxSize-1 {
		return
	}

	if maxSize <= len(subSlice) {
		return
	}

	for i := start; i < maxSize; i++ {
		tempSlice := make([]int, len(subSlice))
		copy(tempSlice, subSlice)
		tempSlice = append(tempSlice, originalSlice[i])
		*finalResult = append(*finalResult, tempSlice)
		//fmt.Println(tempSlice)

		if i+1 < maxSize {
			combine(i+1, originalSlice, tempSlice, finalResult)
		}
	}

}
