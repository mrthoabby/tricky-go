package set

import "iter"

// Set implements a generic set using a map.
// The zero value is ready to use.
type Set[T comparable] struct {
	items map[T]struct{}
}

// New creates a new empty set.
func New[T comparable]() Set[T] {
	return Set[T]{items: make(map[T]struct{})}
}

// Add adds an item to the set.
func (instance *Set[T]) Add(item T) {
	if instance.items == nil {
		instance.items = make(map[T]struct{})
	}
	instance.items[item] = struct{}{}
}

// Remove removes an item from the set.
func (instance *Set[T]) Remove(item T) {
	if instance.items == nil {
		return
	}
	delete(instance.items, item)
}

// Contains checks if an item exists in the set.
func (instance Set[T]) Contains(item T) bool {
	_, exists := instance.items[item]
	return exists
}

// ToSlice returns a slice with all values in the set.
func (instance Set[T]) ToSlice() []T {
	values := make([]T, len(instance.items))
	index := 0
	for item := range instance.items {
		values[index] = item
		index++
	}
	return values
}

// Union returns a new set containing elements from both sets.
func (instance Set[T]) Union(other Set[T]) Set[T] {
	union := make(map[T]struct{}, len(instance.items)+len(other.items))
	for item := range instance.items {
		union[item] = struct{}{}
	}
	for item := range other.items {
		union[item] = struct{}{}
	}
	return Set[T]{items: union}
}

// Intersection returns a new set containing only elements present in both sets.
func (instance Set[T]) Intersection(other Set[T]) Set[T] {
	if len(instance.items) > len(other.items) {
		instance, other = other, instance
	}

	intersection := make(map[T]struct{}, len(instance.items))
	for item := range instance.items {
		if _, exists := other.items[item]; exists {
			intersection[item] = struct{}{}
		}
	}
	return Set[T]{items: intersection}
}

// Length returns the number of elements in the set.
func (instance Set[T]) Length() int {
	return len(instance.items)
}

// All returns an iterator over the elements in the set.
// Iteration order is not guaranteed.
func (instance Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range instance.items {
			if !yield(item) {
				return
			}
		}
	}
}
