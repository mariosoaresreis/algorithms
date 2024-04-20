package domain

func LeadersInArray(a []int) []int {
	result := make([]int, 0, 10)
	start := len(a) - 1
	maximum := a[start]
	result = append(result, maximum)

	for i := len(a) - 2; i >= 0; i-- {
		if a[i] > maximum {
			result = append(result, a[i])
			maximum = a[i]
		}
	}

	return result
}
