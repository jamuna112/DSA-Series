package stack

import (
	"reflect"
	"testing"
)

func TestPushOps(t *testing.T) {
	var stack Stack

	stack.Push(5)
	stack.Push("Hello")
	stack.Push(3)
	stack.Push(2.5)

	expected := Stack{5, "Hello", 3, 2.5}

	if !reflect.DeepEqual(stack, expected) {
		t.Errorf("test failed, expected %v but got %v", expected, stack)
	}

}

func TestPopOps(t *testing.T) {
	var stack Stack

	stack.Push(5)
	stack.Push("Hello")
	stack.Push(3)
	stack.Push(2.5)
	stack.Pop()

	expected := Stack{5, "Hello", 3}

	if !reflect.DeepEqual(stack, expected) {
		t.Errorf("test failed, expected %v but got %v", expected, stack)
	}

}

func TestPopStackIfLenIsEmpty(t *testing.T) {
	var stack Stack

	item, ok := stack.Pop()

	if ok != false {
		t.Errorf("test failed, expected %v but got %v", false, ok)
	}

	if item != nil {
		t.Errorf("test failed, expected %v, but got %v", nil, item)
	}

}

func TestPeek(t *testing.T) {
	var stack Stack

	item, ok := stack.Peek()

	if ok != false {
		t.Errorf("test failed, expected %v but got %v", false, ok)
	}

	if item != nil {
		t.Errorf("test failed, expected %v, but got %v", nil, item)
	}

	stack.Push(5)
	stack.Push(3)
	stack.Push(2)

	item, ok = stack.Peek()

	if item != 2 {
		t.Errorf("test failed, expected %v but got %v", 2, item)
	}

	if ok != true {
		t.Errorf("test failed, expected %v but got %v", true, ok)
	}

}

func TestIsEmptyStack(t *testing.T) {
	var stack Stack
	ok := stack.IsEmpty()

	if ok != true {
		t.Errorf("test failed, expected %v but got %v", true, ok)
	}
}
