package stacks

func IsValid(s1 string) bool {
	stack := []rune{}

	for _, c := range s1 {
		if c == '{' || c == '[' || c == '(' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if c == '}' && top != '{' {
				return false
			}

			if c == ']' && top != '[' {
				return false
			}

			if c == ')' && top != '(' {
				return false
			}
		}
	}

	return len(stack) == 0
}
