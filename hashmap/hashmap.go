package hashmap

import (
	"errors"
	"iter"
)

// Map implements a generic hash map using Go's built-in map.
// The zero value is ready to use.
type Map[K comparable, V any] struct {
	items map[K]V
}

// New creates a new empty map.
func New[K comparable, V any]() Map[K, V] {
	return Map[K, V]{items: make(map[K]V)}
}

// NewWithCapacity creates a new empty map with reserved capacity.
// Returns error if capacity is negative.
func NewWithCapacity[K comparable, V any](capacity int) (Map[K, V], error) {
	if capacity < 0 {
		return Map[K, V]{}, errors.New("capacity must be non-negative")
	}
	return Map[K, V]{items: make(map[K]V, capacity)}, nil
}

// Set stores the value for the given key.
func (instance *Map[K, V]) Set(key K, value V) {
	if instance.items == nil {
		instance.items = make(map[K]V)
	}
	instance.items[key] = value
}

// Get returns the value for the key along with whether it exists.
func (instance Map[K, V]) Get(key K) (V, bool) {
	value, exists := instance.items[key]
	return value, exists
}

// Delete removes the key and its value from the map.
func (instance *Map[K, V]) Delete(key K) {
	if instance.items == nil {
		return
	}
	delete(instance.items, key)
}

// Contains reports whether the key exists in the map.
func (instance Map[K, V]) Contains(key K) bool {
	_, exists := instance.items[key]
	return exists
}

// Len returns the number of items in the map.
func (instance Map[K, V]) Len() int {
	return len(instance.items)
}

// Clear removes all items from the map.
func (instance *Map[K, V]) Clear() {
	instance.items = make(map[K]V)
}

// Keys returns a slice with all keys in the map.
// Iteration order is not guaranteed.
func (instance Map[K, V]) Keys() []K {
	keys := make([]K, 0, len(instance.items))
	for key := range instance.items {
		keys = append(keys, key)
	}
	return keys
}

// Values returns a slice with all values in the map.
// Iteration order is not guaranteed.
func (instance Map[K, V]) Values() []V {
	values := make([]V, 0, len(instance.items))
	for _, value := range instance.items {
		values = append(values, value)
	}
	return values
}

// AsMap returns a shallow copy of the underlying map.
func (instance Map[K, V]) AsMap() map[K]V {
	clone := make(map[K]V, len(instance.items))
	for key, value := range instance.items {
		clone[key] = value
	}
	return clone
}

// All returns an iterator over key/value pairs.
// Iteration order is not guaranteed.
func (instance Map[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for key, value := range instance.items {
			if !yield(key, value) {
				return
			}
		}
	}
}
