package tdd

import (
	"strings"
	"testing"
)

func Test_Wrap(t *testing.T) {
	if !assertWrap("hello", 5, "hello") {
		t.Error("failed!")
	}
}

func Test_Wrap2(t *testing.T) {
	if !assertWrap("hello there", 5, "hello\nthere") {
		t.Error("failed!")
	}
}

func Test_Wrap3(t *testing.T) {
	if !assertWrap("hello there mom", 5, "hello\nthere\nmom") {
		t.Error("failed!")
	}
}

func Test_Wrap4(t *testing.T) {
	println("Four score and seven years ago our fathers brought\nforth on this continent, a new nation, conceived in\nLiberty, and dedicated to the proposition that all\nmen are created equal")
	if !assertWrap(
		"Four score and seven years ago our fathers brought forth on this continent, a new nation, conceived in Liberty, and dedicated to the proposition that all men are created equal",
		52, "Four score and seven years ago our fathers brought\nforth on this continent, a new nation, conceived in\nLiberty, and dedicated to the proposition that all\nmen are created equal") {
		t.Error("failed!")
	}
}

func assertWrap(s string, i int, s2 string) bool {
	return wrap(s, i) == s2
}

func wrap(word string, w int) string {
	if len(word) <= w {
		return word
	}

	newString := word[:w]
	maxIndex := strings.LastIndex(newString, " ")

	if maxIndex != -1 {
		newString = newString[:maxIndex]
	} else {
		maxIndex = w
	}

	result := strings.Trim(newString, " ") + "\n" + wrap(word[maxIndex+1:], w)
	return result
}

func assertWrap2(s string, i int, s2 string) bool {
	return wrap2(s, i) == s2
}

func Test_Wrap5(t *testing.T) {
	// Test A
	if !assertWrap2("hello my little friend", 24, "hello my little friend") {
		t.Error("Test A failed")
	}

	// Test B
	if !assertWrap2("hello my little friend", 17, "hello my...") {
		t.Error("Test B failed")
	}

	// Test C
	if !assertWrap2("hello my little friend", 4, "...") {
		t.Error("Test C failed")
	}

	// Test D
	if !assertWrap2("hello my little friend", 18, "hello my little...") {
		t.Error("Test D failed")
	}

}

func Test_Wrap6(t *testing.T) {
	if !solution([]int{10, 1, 2, 3, 4}) {
		t.Error("Test A failed")
	}

	if solution([]int{1, 2, 1, 2, 3}) == true {
		t.Error("Test B failed")
	}

	if !solution([]int{1, 2, 3, 5, 4, 5}) {
		t.Error("Test C failed")
	}

	if !solution([]int{1, 2, 3, 6, 5}) {
		t.Error("Test D failed")
	}

	if solution([]int{1, 2, 5, 5, 5}) {
		t.Error("Test E failed")
	}

}

func solution(sequence []int) bool {
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

			if len(sequence) >= 3 && i >= 1 && sequence[i] == sequence[i-1] && sequence[i] == sequence[i+1] {
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

func wrap2(word string, w int) string {
	if len(word) <= w {
		return word
	}

	wrappedWord := word

	for len(wrappedWord)+3 > w {
		lastIndex := strings.LastIndex(wrappedWord, " ")

		if lastIndex == -1 {
			return "..."
		}

		wrappedWord = wrappedWord[:lastIndex]
	}

	return wrappedWord + "..."
}
