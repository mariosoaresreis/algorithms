package domain

import "fmt"

func Permute(slice []int, l int, r int, count *int) {
	if l == r {
		*count++
		fmt.Println(slice)
	} else {
		for i := l; i <= r; i++ {
			swap(slice, l, i)
			Permute(slice, l+1, r, count)
			swap(slice, l, i)
		}
	}
}

func swap(slice []int, source int, dest int) {
	//copySlice := make([]int, len(slice))
	//copy(copySlice, slice)
	temp := slice[source]
	slice[source] = slice[dest]
	slice[dest] = temp
}
