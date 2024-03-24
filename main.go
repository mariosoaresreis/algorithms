package main

import (
	"algorithms/com/marioreis/domain"
	"fmt"
)

func main() {
	array := []int{2, 2, 3, 4, 0, 0, 0}
	array2 := []int{1, 2, 4}
	domain.Merge(array, 7, array2, 5)

	println(array)
}

// findElementsWithSum of k from arr of size
func findElementsWithSum(arr [10]int, combinations [19]int, size int, k int, addValue int, l int, m int) int {
	var num int = 0
	if addValue > k {
		return -1
	}
	if addValue == k {
		num = num + 1
		var p int = 0
		for p = 0; p < m; p++ {
			fmt.Printf("%d,", arr[combinations[p]])
		}
		fmt.Println(" ")
	}
	var i int
	for i = l; i < size; i++ {
		combinations[m] = l
		findElementsWithSum(arr, combinations, size, k, addValue+arr[i], l, m+1)
		l = l + 1
	}
	return num
}
