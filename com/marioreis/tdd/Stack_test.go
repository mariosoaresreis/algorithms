package tdd

import "testing"

func Test_CanAddElement(t *testing.T) {
	stack := new(Stack)
	pushOK := stack.Push(3)

	if !pushOK && stack.size == 1 {
		t.Errorf("Can't add element to stack")
	}
}

func Test_CanRemoveElement(t *testing.T) {
	stack := new(Stack)
	stack.Push(3)
	number, err := stack.Pop()

	if err != nil && number != 3 && stack.size == 0 {
		t.Errorf("Can't pop element from stack")
	}
}

func Test_AddAndRemoveElement(t *testing.T) {
	stack := new(Stack)
	stack.Push(3)
	stack.Push(4)
	number1, err1 := stack.Pop()
	number2, err2 := stack.Pop()

	if err1 != nil || err2 != nil {
		t.Errorf("Can't pop stack!")
	}

	if !(number1 == 4 && number2 == 3 && stack.size == 0) {
		t.Errorf("Can't pop and remove element from stack")
	}
}

func Test_RemoveFromEmptyStack(t *testing.T) {
	stack := new(Stack)
	_, err := stack.Pop()

	if err == nil {
		t.Errorf("To pop an empty stack shoud return an error")
	}
}
