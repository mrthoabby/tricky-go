package heap

import "testing"

type item struct {
	priority int
	value    string
}

// TestMinHeap_Push_And_Pop_MaintainsAscendingOrder verifies that a MinHeap
// returns items in ascending order of priority, regardless of insertion order.
func TestMinHeap_Push_And_Pop_MaintainsAscendingOrder(t *testing.T) {
	tests := []struct {
		name  string
		input []item
		want  []item
	}{
		{
			name: "Should sort mixed input ascending",
			input: []item{
				{priority: 3, value: "three"},
				{priority: 1, value: "one"},
				{priority: 2, value: "two"},
			},
			want: []item{
				{priority: 1, value: "one"},
				{priority: 2, value: "two"},
				{priority: 3, value: "three"},
			},
		},
		{
			name: "Should handle already sorted input",
			input: []item{
				{priority: 1, value: "one"},
				{priority: 2, value: "two"},
				{priority: 3, value: "three"},
			},
			want: []item{
				{priority: 1, value: "one"},
				{priority: 2, value: "two"},
				{priority: 3, value: "three"},
			},
		},
		{
			name: "Should handle reverse sorted input",
			input: []item{
				{priority: 3, value: "three"},
				{priority: 2, value: "two"},
				{priority: 1, value: "one"},
			},
			want: []item{
				{priority: 1, value: "one"},
				{priority: 2, value: "two"},
				{priority: 3, value: "three"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange - MinHeap comparator
			h := New(func(a, b item) bool {
				return a.priority < b.priority
			})

			// Act
			for _, v := range tt.input {
				h.Push(v)
			}

			// Assert
			for _, wantItem := range tt.want {
				gotItem, err := h.Pop()
				if err != nil {
					t.Errorf("Pop() returned error: %v", err)
					continue
				}
				if gotItem.priority != wantItem.priority {
					t.Errorf("Pop() got priority %d, want %d", gotItem.priority, wantItem.priority)
				}
			}
		})
	}
}

// TestMaxHeap_Push_And_Pop_MaintainsDescendingOrder verifies that a MaxHeap
// returns items in descending order of priority.
func TestMaxHeap_Push_And_Pop_MaintainsDescendingOrder(t *testing.T) {
	tests := []struct {
		name  string
		input []item
		want  []item
	}{
		{
			name: "Should sort mixed input descending",
			input: []item{
				{priority: 1, value: "one"},
				{priority: 3, value: "three"},
				{priority: 2, value: "two"},
			},
			want: []item{
				{priority: 3, value: "three"},
				{priority: 2, value: "two"},
				{priority: 1, value: "one"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange - MaxHeap comparator
			h := New(func(a, b item) bool {
				return a.priority > b.priority
			})

			// Act
			for _, v := range tt.input {
				h.Push(v)
			}

			// Assert
			for _, wantItem := range tt.want {
				gotItem, err := h.Pop()
				if err != nil {
					t.Errorf("Pop() returned error: %v", err)
					continue
				}
				if gotItem.priority != wantItem.priority {
					t.Errorf("Pop() got priority %d, want %d", gotItem.priority, wantItem.priority)
				}
			}
		})
	}
}

// TestMinHeap_Peek_ReturnsMinElement verifies that Peek returns the minimum element
// without removing it from the heap.
func TestMinHeap_Peek_ReturnsMinElement(t *testing.T) {
	// Arrange
	h := New(func(a, b item) bool {
		return a.priority < b.priority
	})
	input := []item{
		{priority: 10, value: "ten"},
		{priority: 5, value: "five"},
		{priority: 20, value: "twenty"},
	}
	expected := item{priority: 5, value: "five"}

	for _, v := range input {
		h.Push(v)
	}

	// Act
	got, err := h.Peek()
	if err != nil {
		t.Fatalf("Peek() returned error: %v", err)
	}

	// Assert
	if got.priority != expected.priority {
		t.Errorf("Peek() got priority %d, want %d", got.priority, expected.priority)
	}
	if h.Len() != len(input) {
		t.Errorf("Peek() modified heap length, got %d, want %d", h.Len(), len(input))
	}
}

// TestMaxHeap_Peek_ReturnsMaxElement verifies that Peek returns the maximum element
// without removing it from the heap.
func TestMaxHeap_Peek_ReturnsMaxElement(t *testing.T) {
	// Arrange
	h := New(func(a, b item) bool {
		return a.priority > b.priority
	})
	input := []item{
		{priority: 10, value: "ten"},
		{priority: 5, value: "five"},
		{priority: 20, value: "twenty"},
	}
	expected := item{priority: 20, value: "twenty"}

	for _, v := range input {
		h.Push(v)
	}

	// Act
	got, err := h.Peek()
	if err != nil {
		t.Fatalf("Peek() returned error: %v", err)
	}

	// Assert
	if got.priority != expected.priority {
		t.Errorf("Peek() got priority %d, want %d", got.priority, expected.priority)
	}
	if h.Len() != len(input) {
		t.Errorf("Peek() modified heap length, got %d, want %d", h.Len(), len(input))
	}
}

// TestContainer_Len_ReturnsCount verifies the Len method reflects the number of items.
func TestContainer_Len_ReturnsCount(t *testing.T) {
	// Arrange
	h := New(func(a, b item) bool {
		return a.priority < b.priority
	})

	// Assert Initial
	if h.Len() != 0 {
		t.Errorf("New heap should have Len 0, got %d", h.Len())
	}

	// Act & Assert Push
	h.Push(item{priority: 1})
	if h.Len() != 1 {
		t.Errorf("Heap len should be 1 after push, got %d", h.Len())
	}

	h.Push(item{priority: 2})
	if h.Len() != 2 {
		t.Errorf("Heap len should be 2 after push, got %d", h.Len())
	}

	// Act & Assert Pop
	_, err := h.Pop()
	if err != nil {
		t.Fatalf("Pop() returned error: %v", err)
	}
	if h.Len() != 1 {
		t.Errorf("Heap len should be 1 after pop, got %d", h.Len())
	}
}

// TestContainer_Peek_OnEmptyHeap_ReturnsError verifies that calling Peek on an empty heap returns an error.
func TestContainer_Peek_OnEmptyHeap_ReturnsError(t *testing.T) {
	// Arrange
	h := New(func(a, b item) bool {
		return a.priority < b.priority
	})

	// Act
	_, err := h.Peek()

	// Assert
	if err == nil {
		t.Errorf("Peek() on empty heap should return error, got nil")
	}
}

// TestContainer_Pop_OnEmptyHeap_ReturnsError verifies that calling Pop on an empty heap returns an error.
func TestContainer_Pop_OnEmptyHeap_ReturnsError(t *testing.T) {
	// Arrange
	h := New(func(a, b item) bool {
		return a.priority < b.priority
	})

	// Act
	_, err := h.Pop()

	// Assert
	if err == nil {
		t.Errorf("Pop() on empty heap should return error, got nil")
	}
}

// TestPrimitiveTypes verifies that the heap works with primitive types like int.
func TestPrimitiveTypes(t *testing.T) {
	// MinHeap for ints
	h := New(func(a, b int) bool {
		return a < b
	})

	h.Push(3)
	h.Push(1)
	h.Push(2)

	val, err := h.Pop()
	if err != nil || val != 1 {
		t.Errorf("Expected 1, got %d, err: %v", val, err)
	}

	val, err = h.Pop()
	if err != nil || val != 2 {
		t.Errorf("Expected 2, got %d, err: %v", val, err)
	}

	val, err = h.Pop()
	if err != nil || val != 3 {
		t.Errorf("Expected 3, got %d, err: %v", val, err)
	}
}

// TestHeap_Clear_RemovesAllElements verifies clearing the heap.
func TestHeap_Clear_RemovesAllElements(t *testing.T) {
	// Arrange
	h := New(func(a, b int) bool {
		return a < b
	})
	h.Push(1)
	h.Push(2)

	// Act
	h.Clear()

	// Assert
	if h.Len() != 0 {
		t.Errorf("Len should be 0 after clear, got %d", h.Len())
	}
	_, err := h.Pop()
	if err == nil {
		t.Error("Pop should return error after clear")
	}
}
