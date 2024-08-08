package stacks

func NextLargerElement(arr []int) []int {
	res := make([]int, len(arr))
	stack := make([]int, 0)

	for i := len(arr) - 1; i >= 0; i-- {
		if len(stack) == 0 {
			stack = append(stack, arr[i])
			res[i] = -1
		} else {
			for len(stack) > 0 && arr[i] >= stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}

			if len(stack) > 0 {
				res[i] = stack[len(stack)-1]
			} else {
				res[i] = -1
			}

			stack = append(stack, arr[i])
		}
	}

	return res
}
