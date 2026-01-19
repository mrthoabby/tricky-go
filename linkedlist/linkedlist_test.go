package linkedlist

import (
	"reflect"
	"testing"
)

// TestNew_CreatesEmptyList verifies that a new list is initialized correctly.
func TestNew_CreatesEmptyList(t *testing.T) {
	// Act
	l := New[int]()

	// Assert
	if l.Len() != 0 {
		t.Errorf("New list should be empty, got len %d", l.Len())
	}
	if l.Front() != nil {
		t.Error("New list Front() should be nil")
	}
	if l.Back() != nil {
		t.Error("New list Back() should be nil")
	}
}

// TestList_PushBack_AddsElementToEnd verifies adding elements to the back.
func TestList_PushBack_AddsElementToEnd(t *testing.T) {
	// Arrange
	l := New[int]()

	// Act
	l.PushBack(1)
	l.PushBack(2)

	// Assert
	if l.Len() != 2 {
		t.Errorf("List len should be 2, got %d", l.Len())
	}
	if l.Back().Value != 2 {
		t.Errorf("Back element should be 2, got %d", l.Back().Value)
	}
	if l.Front().Value != 1 {
		t.Errorf("Front element should be 1, got %d", l.Front().Value)
	}
}

// TestList_PushFront_AddsElementToStart verifies adding elements to the front.
func TestList_PushFront_AddsElementToStart(t *testing.T) {
	// Arrange
	l := New[int]()

	// Act
	l.PushFront(1)
	l.PushFront(2)

	// Assert
	if l.Len() != 2 {
		t.Errorf("List len should be 2, got %d", l.Len())
	}
	if l.Front().Value != 2 {
		t.Errorf("Front element should be 2, got %d", l.Front().Value)
	}
	if l.Back().Value != 1 {
		t.Errorf("Back element should be 1, got %d", l.Back().Value)
	}
}

// TestList_Remove_RemovesCorrectElement verifies removing elements.
func TestList_Remove_RemovesCorrectElement(t *testing.T) {
	// Arrange
	l := New[int]()
	l.PushBack(1)
	middleNode := l.PushBack(2)
	l.PushBack(3)

	// Act
	_, err := l.Remove(middleNode)
	if err != nil {
		t.Fatalf("Remove() returned error: %v", err)
	}

	// Assert
	expected := []int{1, 3}
	if !reflect.DeepEqual(l.ToSlice(), expected) {
		t.Errorf("Expected list %v, got %v", expected, l.ToSlice())
	}
	if l.Len() != 2 {
		t.Errorf("List len should be 2, got %d", l.Len())
	}
}

// TestList_Remove_HeadAndTail verifies removing head and tail specifically.
func TestList_Remove_HeadAndTail(t *testing.T) {
	// Arrange
	l := New[int]()
	l.PushBack(1)
	l.PushBack(2)

	// Act - Remove Head
	_, err := l.Remove(l.Front())
	if err != nil {
		t.Fatalf("Remove() returned error: %v", err)
	}

	// Assert - Remove Head
	if l.Len() != 1 {
		t.Errorf("Len should be 1 after removing head, got %d", l.Len())
	}
	if l.Front().Value != 2 {
		t.Errorf("New head should be 2, got %d", l.Front().Value)
	}

	// Act - Remove Tail
	_, err = l.Remove(l.Back())
	if err != nil {
		t.Fatalf("Remove() returned error: %v", err)
	}

	// Assert - Remove Tail
	if l.Len() != 0 {
		t.Errorf("Len should be 0 after removing tail, got %d", l.Len())
	}
	if l.Front() != nil {
		t.Error("Front should be nil on empty list")
	}
}

// TestList_ToSlice_ReturnsAllElementsInOrder verifies slice conversion.
func TestList_ToSlice_ReturnsAllElementsInOrder(t *testing.T) {
	// Arrange
	l := New[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	// Act
	got := l.ToSlice()

	// Assert
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

// TestList_Remove_NilNode_ReturnsError verifies that removing a nil node returns an error.
func TestList_Remove_NilNode_ReturnsError(t *testing.T) {
	// Arrange
	l := New[int]()
	l.PushBack(1)

	// Act
	_, err := l.Remove(nil)

	// Assert
	if err == nil {
		t.Error("Remove() with nil node should return error, got nil")
	}
}

// TestList_Clear_RemovesAllElements verifies clearing the list.
func TestList_Clear_RemovesAllElements(t *testing.T) {
	// Arrange
	l := New[int]()
	l.PushBack(1)
	l.PushBack(2)

	// Act
	l.Clear()

	// Assert
	if l.Len() != 0 {
		t.Errorf("List len should be 0 after clear, got %d", l.Len())
	}
	if l.Front() != nil {
		t.Error("Front should be nil after clear")
	}
	if l.Back() != nil {
		t.Error("Back should be nil after clear")
	}
}

// TestList_All_IteratesElements verifies the iterator.
func TestList_All_IteratesElements(t *testing.T) {
	// Arrange
	l := New[int]()
	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)

	// Act
	var got []int
	for v := range l.All() {
		got = append(got, v)
	}

	// Assert
	expected := []int{10, 20, 30}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
