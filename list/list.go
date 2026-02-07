package list

import (
	"errors"
	"iter"
	"sort"
)

// List implements a generic, slice-backed list.
// It is not safe for concurrent use.
type List[T any] struct {
	items []T
}

// New creates a new empty list.
func New[T any]() *List[T] {
	return &List[T]{items: make([]T, 0)}
}

// NewWithCapacity creates an empty list with reserved capacity.
// Returns error if capacity is negative.
func NewWithCapacity[T any](capacity int) (*List[T], error) {
	if capacity < 0 {
		return nil, errors.New("capacity must be non-negative")
	}
	return &List[T]{items: make([]T, 0, capacity)}, nil
}

// Add appends an item to the list.
func (instance *List[T]) Add(item T) {
	instance.items = append(instance.items, item)
}

// Remove deletes the item at the given index and returns it.
// Returns error if the index is out of range.
func (instance *List[T]) Remove(index int) (T, error) {
	if index < 0 || index >= len(instance.items) {
		var zero T
		return zero, errors.New("index out of range")
	}
	item := instance.items[index]
	copy(instance.items[index:], instance.items[index+1:])
	lastIndex := len(instance.items) - 1
	var zero T
	instance.items[lastIndex] = zero
	instance.items = instance.items[:lastIndex]
	return item, nil
}

// Sort orders the list in-place using the provided comparator.
// Returns error if the comparator is nil.
func (instance *List[T]) Sort(less func(a, b T) bool) error {
	if less == nil {
		return errors.New("less function is nil")
	}
	sort.Slice(instance.items, func(i, j int) bool {
		return less(instance.items[i], instance.items[j])
	})
	return nil
}

// Get returns the item at the given index.
// Returns error if the index is out of range.
func (instance *List[T]) Get(index int) (T, error) {
	if index < 0 || index >= len(instance.items) {
		var zero T
		return zero, errors.New("index out of range")
	}
	return instance.items[index], nil
}

// Len returns the number of elements in the list.
func (instance *List[T]) Len() int {
	return len(instance.items)
}

// Clear removes all elements from the list.
func (instance *List[T]) Clear() {
	instance.items = make([]T, 0)
}

// ToSlice returns a copy of the list items.
func (instance *List[T]) ToSlice() []T {
	items := make([]T, len(instance.items))
	copy(items, instance.items)
	return items
}

// All returns an iterator over the elements in the list.
func (instance *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, item := range instance.items {
			if !yield(item) {
				return
			}
		}
	}
}
