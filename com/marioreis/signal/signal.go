package signal

/*
Given a rectangular matrix of characters, add a border of asterisks(*) to it.

Example

For

picture = ["abc",
           "ded"]
the output should be

solution(picture) = ["*****",
                     "*abc*",
                     "*ded*",
                     "*****"]
*/

func CodeWriting(picture []string) []string {
	result := make([]string, 0)
	asterix := "*"

	for i := 0; i < len(picture[0]); i++ {
		asterix += "*"
	}

	asterix += "*"

	result = append(result, asterix)

	for _, v := range picture {
		result = append(result, "*"+v+"*")
	}

	result = append(result, asterix)
	println(result)

	return result
}
