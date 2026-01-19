package set

import (
	"sort"
	"testing"
)

func TestSet(t *testing.T) {
	s := New[int]()

	if s.Length() != 0 {
		t.Errorf("expected empty set, got length %d", s.Length())
	}

	s.Add(1)
	s.Add(2)
	s.Add(2) // duplicate

	if s.Length() != 2 {
		t.Errorf("expected length 2, got %d", s.Length())
	}

	if !s.Contains(1) || !s.Contains(2) {
		t.Error("set should contain 1 and 2")
	}

	if s.Contains(3) {
		t.Error("set should not contain 3")
	}

	s.Remove(1)
	if s.Contains(1) {
		t.Error("set should not contain 1 after removal")
	}

	// Test Union
	s2 := New[int]()
	s2.Add(2)
	s2.Add(3)

	union := s.Union(s2)
	if union.Length() != 2 { // 2, 3
		t.Errorf("expected union length 2, got %d", union.Length())
	}
	if !union.Contains(2) || !union.Contains(3) {
		t.Error("union should contain 2 and 3")
	}

	// Test Intersection
	s3 := New[int]()
	s3.Add(3)
	s3.Add(4)

	s4 := New[int]()
	s4.Add(2)
	s4.Add(3)

	intersection := s3.Intersection(s4)
	if intersection.Length() != 1 {
		t.Errorf("expected intersection length 1, got %d", intersection.Length())
	}
	if !intersection.Contains(3) {
		t.Error("intersection should contain 3")
	}

	// Test ToSlice
	slice := s4.ToSlice()
	sort.Ints(slice)
	if len(slice) != 2 || slice[0] != 2 || slice[1] != 3 {
		t.Errorf("unexpected slice content: %v", slice)
	}
}

func TestSet_All_IteratesElements(t *testing.T) {
	s := New[int]()
	s.Add(10)
	s.Add(20)
	s.Add(30)

	var got []int
	for v := range s.All() {
		got = append(got, v)
	}
	sort.Ints(got)

	if len(got) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(got))
	}
	if got[0] != 10 || got[1] != 20 || got[2] != 30 {
		t.Errorf("Expected [10 20 30], got %v", got)
	}
}

func TestZeroValueUsable(t *testing.T) {
	var s Set[int]

	s.Add(1)
	if !s.Contains(1) {
		t.Fatalf("expected set to contain added item")
	}

	s.Remove(1)
	if s.Contains(1) {
		t.Fatalf("expected item to be removed")
	}
}
