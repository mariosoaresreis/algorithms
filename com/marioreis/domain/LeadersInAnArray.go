package domain

type Stack []int

func (s *Stack) push(e int) {
	*s = append(*s, e)
}

func (s *Stack) pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func LeadersInArray(a []int) []int {
	stack := Stack{}
	start := len(a) - 1
	maximum := a[start]
	stack.push(maximum)

	for i := len(a) - 2; i >= 0; i-- {
		if a[i] > maximum {
			stack.push(a[i])
			maximum = a[i]
		}
	}

	return stack
}
