package stack_test

import (
	"gen/v5/stack"
	"testing"
)


func TestStack(t *testing.T) {

	t.Run("integer stack", func(t *testing.T) {
		// check if it is empty
		myStackOfInts := new(stack.Stack[int])
		AssertTrue(t, myStackOfInts.IsEmpty())

		// check if it is not empty after an element is added
		myStackOfInts.Push(1)
		AssertFalse(t, myStackOfInts.IsEmpty())

		// add an element, pop it and check
		myStackOfInts.Push(2)
		val, _ := myStackOfInts.Pop()
		AssertEqual[int](t, val, 2)
		val, _ = myStackOfInts.Pop()
		AssertEqual[int](t, val, 1)
		AssertTrue(t, myStackOfInts.IsEmpty())
	})
}


func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v want false", got)
	}
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("got %v want %v", got, want)
	}
}