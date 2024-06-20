package tdd

func sortByHeight(a []int) []int {
	for j := 0; j < len(a)-1; j++ {
		for i := len(a) - 1; i >= 0; i-- {
			if a[i] != -1 {
				swapB(a, i)
			}
		}

	}

	return a

}

func swapB(a []int, source int) {
	if a[source] != -1 {
		nextPos := findNextToSwapBack(a, source)

		if nextPos != -1 && a[nextPos] > a[source] {
			temp := a[nextPos]
			a[nextPos] = a[source]
			a[source] = temp
		}

	}
}

func findNextToSwapBack(a []int, source int) int {
	if source >= 1 {
		for i := source - 1; i >= 0; i-- {
			if a[i] != -1 {
				return i
			}
		}
	}

	return -1
}
