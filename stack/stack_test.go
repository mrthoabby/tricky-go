package stack

import "testing"

func TestStack(t *testing.T) {
	s := New[int]()

	if !s.IsEmpty() {
		t.Error("New stack should be empty")
	}

	s.Push(10)
	s.Push(20)

	if s.Len() != 2 {
		t.Errorf("Expected length 2, got %d", s.Len())
	}

	val, err := s.Peek()
	if err != nil || val != 20 {
		t.Errorf("Expected peek 20, got %d, err: %v", val, err)
	}

	val, err = s.Pop()
	if err != nil || val != 20 {
		t.Errorf("Expected pop 20, got %d, err: %v", val, err)
	}

	val, err = s.Pop()
	if err != nil || val != 10 {
		t.Errorf("Expected pop 10, got %d, err: %v", val, err)
	}

	if !s.IsEmpty() {
		t.Error("Stack should be empty after popping all elements")
	}

	_, err = s.Pop()
	if err == nil {
		t.Error("Expected error when popping from empty stack")
	}
}

func TestStack_All_IteratesLIFO(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	var got []int
	for v := range s.All() {
		got = append(got, v)
	}

	// Should iterate from top (3) to bottom (1)
	if len(got) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(got))
	}
	if got[0] != 3 || got[1] != 2 || got[2] != 1 {
		t.Errorf("Expected [3 2 1], got %v", got)
	}
}
