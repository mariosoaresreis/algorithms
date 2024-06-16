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
