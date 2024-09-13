package subsets


func GenerateValidParentheses(num int) []string {
	result := []string{}
	parenthesesString := make([]rune, num*2)
	GenerateParentheses(num, 0, 0, 0, parenthesesString, &result)
	return result
}

func GenerateParentheses(num int, open int, close int, index int, parenthesesString []rune,
	result *[]string){
	if open == num && close == num {
		*result = append(*result, string(parenthesesString))
	}else {
		if open < num {
			parenthesesString[index] = '('
			GenerateParentheses(num, open + 1, close, index+1, parenthesesString, result)
		}

		if close < open {
			parenthesesString[index] = ')'
			GenerateParentheses(num, open, close + 1, index+1, parenthesesString, result)
		}
	}
