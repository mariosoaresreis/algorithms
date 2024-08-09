package main

import (
	"algorithms/com/marioreis/grokking/stacks"
	"fmt"
)

type Test struct {
	Id *int
}

// deploy
func main() {
	stacks.SimplifyPath("/a//b////c/d//././/..")
}

func printar(t Test) {
	println(fmt.Sprintf("array: %v", t.Id))
}

func findSumOfDigits(num int) int {
	if num == 0 {
		return 0
	}

	sum := 0
	var digit int

	for num > 0 {
		digit = num % 10
		sum += digit
		num /= 10
	}

	return sum
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
