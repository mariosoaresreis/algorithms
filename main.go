package main

import (
	"algorithms/com/marioreis/domain"
	"fmt"
)

func main() {
	count := 0
	array := []int{1, 2, 3, 4}
	domain.Permute(array, 0, len(array)-1, &count)
	fmt.Println(count)
}
