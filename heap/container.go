package heap

import (
	"errors"

	goheap "container/heap"
)

// Container wraps our Heap implementation to provide a simpler API
// using container/heap internally.
type Container[T any] struct {
	internal *heap[T]
}

// New creates and initializes a new heap container.
// less is a function that returns true if a should be sorted before b.
// For a MinHeap, use func(a, b T) bool { return a < b }.
// For a MaxHeap, use func(a, b T) bool { return a > b }.
// less must not be nil.
func New[T any](less func(a, b T) bool) *Container[T] {
	if less == nil {
		panic("heap.New: less comparator must not be nil")
	}
	heap := &heap[T]{
		items: make([]T, 0),
		less:  less,
	}
	goheap.Init(heap)
	return &Container[T]{
		internal: heap,
	}
}

// Push adds an item to the heap maintaining the heap property.
func (instance *Container[T]) Push(item T) {
	goheap.Push(instance.internal, item)
}

// Pop removes and returns the top priority item from the heap.
// Returns error if heap is empty.
func (instance *Container[T]) Pop() (T, error) {
	if instance.Len() == 0 {
		var zero T
		return zero, errors.New("heap is empty")
	}
	return goheap.Pop(instance.internal).(T), nil
}

// Peek returns the top priority item without removing it.
// Returns error if heap is empty.
func (instance *Container[T]) Peek() (T, error) {
	if instance.Len() == 0 {
		var zero T
		return zero, errors.New("heap is empty")
	}
	return instance.internal.items[0], nil
}

// Len returns the number of items in the heap.
func (instance *Container[T]) Len() int {
	return instance.internal.Len()
}

// Clear removes all elements from the heap.
// It does not re-heapify; it just drops all items.
func (instance *Container[T]) Clear() {
	instance.internal.items = make([]T, 0)
}
