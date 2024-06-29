package signal

import (
	"fmt"
	"slices"
	"strings"
	"testing"
	"unicode"
)

func Test_reverse(t *testing.T) {
	inputString := "(1(23(45)6(78)9))"
	//word2 := "123556879"

	for strings.LastIndex(inputString, ")") != -1 {
		println(inputString)
		begin := strings.LastIndex(inputString, "(")
		println(begin)
		end := lastIndex(inputString, begin, len(inputString)-1)
		println(end)
		inputString = reverse([]byte(inputString), begin, end)
	}

	println(fmt.Sprintf("final result %s", inputString))
}

func lastIndex(word string, begin int, end int) int {
	for i := begin; i <= end; i++ {
		if word[i] == ')' {
			return i
		}
	}

	return -1
}

func reverse(word []byte, begin int, end int) string {
	j := 0
	e := (end-begin)/2 + begin

	if end%2 == 0 {
		e = (end-begin)/2 + begin
	}

	for i := begin; i <= e; i++ {
		temp := word[i]
		word[i] = word[end-j]
		word[end-j] = temp
		j++
	}
	result := slices.Concat(slices.Concat(word[0:begin], word[begin+1:end]), word[end+1:])
	r := string(result)
	return r
}

func Test_almostIncreasingArray(t *testing.T) {
	/*if !almostIncreasingArray([]int{10, 1, 2, 3, 4}) {
		t.Error("Test A failed")
	}

	if almostIncreasingArray([]int{1, 2, 1, 2, 3}) {
		t.Error("Test B failed")
	}

	if !almostIncreasingArray([]int{1, 2, 3, 5, 4, 5}) {
		t.Error("Test C failed")
	}

	if !almostIncreasingArray([]int{1, 2, 3, 6, 5}) {
		t.Error("Test D failed")
	}

	if almostIncreasingArray([]int{1, 2, 5, 5, 5}) {
		t.Error("Test E failed")
	}

	if !almostIncreasingArray([]int{1, 2, 3, 4, 5, 3}) {
		t.Error("Test F failed")
	}

	if almostIncreasingArray([]int{1, 2, 18, 4, 4}) {
		t.Error("Test G failed")
	}*/

	//s := "1name"
	var balance float32 = 100
	denominador := float32(float32(20) / float32(100))
	percentage := float32(balance * denominador)
	balance += percentage
	println(balance)

}

func testVariable(name string) bool {
	if !unicode.IsLetter(rune(name[0])) || unicode.IsDigit(rune(name[0])) || rune(name[0]) == ' ' || unicode.IsSymbol(rune(name[0])) {
		return false
	}

	if len(name) >= 1 {
		for i := 1; i < len(name); i++ {
			if (name[i] != '_') &&
				(!unicode.IsLetter(rune(name[i])) && !unicode.IsDigit(rune(name[i]))) {
				return false
			}
		}

	} else {
		if !unicode.IsLetter(rune(name[0])) {
			return false
		}
	}

	return true

}

func almostIncreasingArray(sequence []int) bool {
	lastNum := sequence[0]
	count := 0

	if len(sequence) <= 2 {
		return true
	}

	for i := 1; i < len(sequence); i++ {
		if sequence[i] <= lastNum {
			count++

			if i == 1 {
				lastNum = sequence[i]
				continue
			}

			if i < len(sequence)-1 && len(sequence) >= 3 && sequence[i] <= sequence[i-1] && sequence[i+1] <= sequence[i] {
				return false
			}

			if i <= len(sequence)-2 && len(sequence) >= 3 && i >= 1 && sequence[i] == sequence[i-1] &&
				sequence[i] == sequence[i+1] {
				return false
			}

			if i <= len(sequence)-2 && sequence[i-1] >= sequence[i+1] {
				if i >= 2 && sequence[i] <= sequence[i-2] {
					return false
				}

				if i >= 2 {
					lastNum = sequence[i-2]
					continue
				} else {
					lastNum = sequence[i-1]
					continue
				}

			}

			if i <= len(sequence)-2 && i >= 1 && sequence[i-1] <= sequence[i+1] {
				lastNum = sequence[i-1]
				continue
			}

		}

		lastNum = sequence[i]
	}

	return count < 2
}
