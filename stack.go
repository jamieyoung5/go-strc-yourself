package strc

import (
	"container/list"
	"sync"
)

// Stack represents a generic stack data structure.
//
// It supports basic operations such as Push, Pop, Peek, IsEmpty, and Size.
//
// Example usage:
//
//	stack := NewStack[int]()
//	stack.Push(1)
//	value := stack.Pop()
//	isEmpty := stack.IsEmpty()
//	size := stack.Size()
type Stack[T any] struct {
	list *list.List
	lock sync.Mutex
}

// NewStack creates a new instance of Stack with the specified type.
// It returns a pointer to the created Stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{list: list.New()}
}

// Push adds a new element to the top of the stack.
func (s *Stack[T]) Push(value T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.list.PushBack(value)
}

// Pop removes and returns the top element from the stack.
// If the stack is empty, it returns the zero value of type T.
func (s *Stack[T]) Pop() T {
	s.lock.Lock()
	defer s.lock.Unlock()

	element := s.list.Back()
	if element != nil {
		s.list.Remove(element)
		return element.Value.(T)
	}

	var zeroValue T
	return zeroValue
}

// Peek returns the top element of the stack without removing it.
// If the stack is empty, it returns the zero value of the element type.
// It locks the stack before accessing and unlocks it afterwards.
func (s *Stack[T]) Peek() T {
	s.lock.Lock()
	defer s.lock.Unlock()

	element := s.list.Back()
	if element != nil {
		return element.Value.(T)
	}

	var zeroValue T
	return zeroValue
}

// IsEmpty returns true if the stack is empty, and false otherwise.
func (s *Stack[T]) IsEmpty() bool {
	return s.list.Len() == 0
}

// Size returns the number of elements in the stack.
// It utilizes the Len method from the underlying linked list.
func (s *Stack[T]) Size() int {
	return s.list.Len()
}
