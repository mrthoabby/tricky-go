package list

import "testing"

func TestListAddRemove(t *testing.T) {
	l := New[int]()
	l.Add(10)
	l.Add(20)
	l.Add(30)

	if l.Len() != 3 {
		t.Fatalf("expected length 3, got %d", l.Len())
	}

	removed, err := l.Remove(1)
	if err != nil {
		t.Fatalf("unexpected error removing item: %v", err)
	}
	if removed != 20 {
		t.Fatalf("expected removed value 20, got %d", removed)
	}

	items := l.ToSlice()
	if len(items) != 2 || items[0] != 10 || items[1] != 30 {
		t.Fatalf("expected [10 30], got %v", items)
	}
}

func TestListRemoveOutOfRange(t *testing.T) {
	l := New[int]()
	l.Add(1)

	_, err := l.Remove(-1)
	if err == nil {
		t.Fatal("expected error for negative index")
	}

	_, err = l.Remove(1)
	if err == nil {
		t.Fatal("expected error for out of range index")
	}
}

func TestListSort(t *testing.T) {
	l := New[int]()
	l.Add(3)
	l.Add(1)
	l.Add(2)

	if err := l.Sort(func(a, b int) bool { return a < b }); err != nil {
		t.Fatalf("unexpected error sorting: %v", err)
	}

	items := l.ToSlice()
	if len(items) != 3 || items[0] != 1 || items[1] != 2 || items[2] != 3 {
		t.Fatalf("expected [1 2 3], got %v", items)
	}
}

func TestListSortNilLess(t *testing.T) {
	l := New[int]()
	if err := l.Sort(nil); err == nil {
		t.Fatal("expected error for nil comparator")
	}
}

func TestListAll(t *testing.T) {
	l := New[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	var got []int
	for v := range l.All() {
		got = append(got, v)
	}

	if len(got) != 3 || got[0] != 1 || got[1] != 2 || got[2] != 3 {
		t.Fatalf("expected [1 2 3], got %v", got)
	}
}

func TestListNewWithCapacity(t *testing.T) {
	l, err := NewWithCapacity[int](5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if l.Len() != 0 {
		t.Fatalf("expected empty list, got %d", l.Len())
	}

	_, err = NewWithCapacity[int](-1)
	if err == nil {
		t.Fatal("expected error for negative capacity")
	}
}
