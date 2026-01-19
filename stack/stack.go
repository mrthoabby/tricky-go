package stack

import (
	"errors"
	"iter"
)

// Stack implements a generic LIFO (Last In, First Out) stack.
// It is not safe for concurrent use.
type Stack[T any] struct {
	items []T
}

// New creates a new empty stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

// Push adds an element to the top of the stack.
func (instance *Stack[T]) Push(item T) {
	instance.items = append(instance.items, item)
}

// Pop removes and returns the element from the top of the stack.
// Returns error if the stack is empty.
func (instance *Stack[T]) Pop() (T, error) {
	if instance.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}
	lastIndex := len(instance.items) - 1
	item := instance.items[lastIndex]
	var zero T
	instance.items[lastIndex] = zero
	instance.items = instance.items[:lastIndex]
	return item, nil
}

// Peek returns the element from the top without removing it.
// Returns error if the stack is empty.
func (instance *Stack[T]) Peek() (T, error) {
	if instance.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}
	return instance.items[len(instance.items)-1], nil
}

// IsEmpty returns true if the stack has no elements.
func (instance *Stack[T]) IsEmpty() bool {
	return len(instance.items) == 0
}

// Len returns the number of elements in the stack.
func (instance *Stack[T]) Len() int {
	return len(instance.items)
}

// Clear empties the stack.
func (instance *Stack[T]) Clear() {
	instance.items = make([]T, 0)
}

// All returns an iterator over the elements in the stack (LIFO order).
func (instance *Stack[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := len(instance.items) - 1; i >= 0; i-- {
			if !yield(instance.items[i]) {
				return
			}
		}
	}
}
