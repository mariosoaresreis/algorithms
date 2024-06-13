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

	pushOK := stack.Push(3)

	if !pushOK && stack.size == 1 {
		t.Errorf("Can't add element to stack")
	}
}

func Test_CanRemoveElement(t *testing.T) {
	stack := &Stack{}

	if stack == nil {
		t.Errorf("Can't create stack")
	}

	stack.Push(3)
	number := stack.Pop()

	if number != 3 && stack.size == 0 {
		t.Errorf("Can't pop element from stack")
	}
}

func Test_AddAndRemoveElement(t *testing.T) {
	stack := &Stack{}

	if stack == nil {
		t.Errorf("Can't create stack")
	}

	stack.Push(3)
	stack.Push(4)
	number1 := stack.Pop()
	number2 := stack.Pop()

	if !(number1 == 4 && number2 == 3 && stack.size == 0) {
		t.Errorf("Can't pop and remove element from stack")
	}
}
