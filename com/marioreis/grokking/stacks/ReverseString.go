package stacks

func ReverseString(sInput string) string {
	result := []rune{}

	for i := len(sInput) - 1; i >= 0; i-- {
		result = append(result, rune(sInput[i]))
	}
	return string(result)
}
