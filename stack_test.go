package strc

import (
	"testing"
)

func TestPush(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
	}

	for _, tt := range tests {
		s := NewStack[int]()
		for _, v := range tt.input {
			s.Push(v)
		}
		for i := len(tt.input) - 1; i >= 0; i-- {
			if s.Peek() != tt.want[i] {
				t.Errorf("Push() = %v, want %v", s.Peek(), tt.want[i])
			}
			s.Pop()
		}
	}
}

func TestPop(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		s := NewStack[int]()
		for _, v := range tt.input {
			s.Push(v)
		}
		for _, v := range tt.want {
			if s.Pop() != v {
				t.Errorf("Pop() = %v, want %v", s.Pop(), v)
			}
		}
		if !s.IsEmpty() {
			t.Errorf("Stack is not empty")
		}
	}
}

func TestPeek(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
		{[]int{1, 2, 3}, 3},
	}

	for _, tt := range tests {
		s := NewStack[int]()
		for _, v := range tt.input {
			s.Push(v)
		}
		if s.Peek() != tt.want {
			t.Errorf("Peek() = %v, want %v", s.Peek(), tt.want)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	s := NewStack[int]()
	if !s.IsEmpty() {
		t.Errorf("IsEmpty() = %v, want %v", s.IsEmpty(), true)
	}
	s.Push(1)
	if s.IsEmpty() {
		t.Errorf("IsEmpty() = %v, want %v", s.IsEmpty(), false)
	}
}

func TestSize(t *testing.T) {
	s := NewStack[int]()
	if s.Size() != 0 {
		t.Errorf("Size() = %v, want %v", s.Size(), 0)
	}
	s.Push(1)
	if s.Size() != 1 {
		t.Errorf("Size() = %v, want %v", s.Size(), 1)
	}
	s.Push(2)
	if s.Size() != 2 {
		t.Errorf("Size() = %v, want %v", s.Size(), 2)
	}
	s.Pop()
	if s.Size() != 1 {
		t.Errorf("Size() = %v, want %v", s.Size(), 1)
	}
}
