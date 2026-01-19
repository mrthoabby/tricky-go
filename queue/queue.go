package queue

import (
	"errors"
	"iter"
)

// Queue implements a generic FIFO (First In, First Out) queue.
// It uses an underlying slice and optimizes memory occasionally.
// It is not safe for concurrent use.
type Queue[T any] struct {
	items []T
}

// New creates a new empty queue.
func New[T any]() *Queue[T] {
	return &Queue[T]{items: make([]T, 0)}
}

// Enqueue adds an item to the end of the queue.
func (instance *Queue[T]) Enqueue(item T) {
	instance.items = append(instance.items, item)
}

// Dequeue removes and returns the first element of the queue.
// Returns error if the queue is empty.
func (instance *Queue[T]) Dequeue() (T, error) {
	if instance.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}
	item := instance.items[0]
	var zero T
	instance.items[0] = zero
	instance.items = instance.items[1:]

	// Optimization: if the underlying slice grows too large and we use little, compact it.
	// This prevents memory leaks in long-lived queues.
	if len(instance.items) > 0 && len(instance.items) < cap(instance.items)/4 {
		newItems := make([]T, len(instance.items))
		copy(newItems, instance.items)
		instance.items = newItems
	}

	return item, nil
}

// Peek returns the first element without removing it.
// Returns error if the queue is empty.
func (instance *Queue[T]) Peek() (T, error) {
	if instance.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}
	return instance.items[0], nil
}

// IsEmpty returns true if the queue has no elements.
func (instance *Queue[T]) IsEmpty() bool {
	return len(instance.items) == 0
}

// Len returns the number of elements in the queue.
func (instance *Queue[T]) Len() int {
	return len(instance.items)
}

// Clear empties the queue.
func (instance *Queue[T]) Clear() {
	instance.items = make([]T, 0)
}

// All returns an iterator over the elements in the queue (FIFO order).
func (instance *Queue[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, item := range instance.items {
			if !yield(item) {
				return
			}
		}
	}
}
