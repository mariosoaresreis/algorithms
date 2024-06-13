package tdd

import "testing"

func Test_CanCreateStack(t *testing.T) {
	stack := &Stack{}

	if stack == nil {
		t.Errorf("Can't create stack")
	}
}

func Test_CanAddElement(t *testing.T) {
	stack := &Stack{}

	if stack == nil {
		t.Errorf("Can't create stack")
	}

	pushOK := stack.push(3)

	if !pushOK && stack.size == 1 {
		t.Errorf("Can't add element to stack")
	}
}

func Test_CanRemoveElement(t *testing.T) {
	stack := &Stack{}

	if stack == nil {
		t.Errorf("Can't create stack")
	}

	stack.push(3)
	number := stack.pop()

	if number != 3 && stack.size == 0 {
		t.Errorf("Can't pop element from stack")
	}
}

func Test_AddAndRemoveElement(t *testing.T) {
	stack := &Stack{}

	if stack == nil {
		t.Errorf("Can't create stack")
	}

	stack.push(3)
	stack.push(4)
	number1 := stack.pop()
	number2 := stack.pop()

	if !(number1 == 4 && number2 == 3 && stack.size == 0) {
		t.Errorf("Can't pop and remove element from stack")
	}
}
