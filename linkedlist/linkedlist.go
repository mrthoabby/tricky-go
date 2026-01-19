package linkedlist

import (
	"errors"
	"iter"
)

// Node represents a node in the linked list.
type Node[T any] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
	list  *List[T]
}

// List implements a doubly linked list.
type List[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// New creates a new empty linked list.
func New[T any]() *List[T] {
	return &List[T]{}
}

// PushBack adds a value to the end of the list.
func (instance *List[T]) PushBack(v T) *Node[T] {
	node := &Node[T]{Value: v, list: instance}
	if instance.tail == nil {
		instance.head = node
		instance.tail = node
	} else {
		node.Prev = instance.tail
		instance.tail.Next = node
		instance.tail = node
	}
	instance.size++
	return node
}

// PushFront adds a value to the start of the list.
func (instance *List[T]) PushFront(v T) *Node[T] {
	node := &Node[T]{Value: v, list: instance}
	if instance.head == nil {
		instance.head = node
		instance.tail = node
	} else {
		node.Next = instance.head
		instance.head.Prev = node
		instance.head = node
	}
	instance.size++
	return node
}

// Front returns the first node of the list or nil if it is empty.
func (instance *List[T]) Front() *Node[T] {
	return instance.head
}

// Back returns the last node of the list or nil if it is empty.
func (instance *List[T]) Back() *Node[T] {
	return instance.tail
}

// Remove removes a node from the list and returns its value.
// Returns error if the node does not belong to this list or is nil.
func (instance *List[T]) Remove(node *Node[T]) (T, error) {
	if node == nil {
		var zero T
		return zero, errors.New("node is nil")
	}
	if node.list != instance {
		var zero T
		return zero, errors.New("node does not belong to this list")
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		instance.head = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		instance.tail = node.Prev
	}
	node.list = nil
	node.Next = nil
	node.Prev = nil
	instance.size--
	return node.Value, nil
}

// Len returns the number of elements.
func (instance *List[T]) Len() int {
	return instance.size
}

// Clear removes all elements from the list.
func (instance *List[T]) Clear() {
	instance.head = nil
	instance.tail = nil
	instance.size = 0
}

// All returns an iterator over the elements in the list.
func (instance *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for n := instance.head; n != nil; n = n.Next {
			if !yield(n.Value) {
				return
			}
		}
	}
}

// ToSlice returns a slice with all values in the list.
func (instance *List[T]) ToSlice() []T {
	items := make([]T, 0, instance.size)
	for n := instance.head; n != nil; n = n.Next {
		items = append(items, n.Value)
	}
	return items
}
