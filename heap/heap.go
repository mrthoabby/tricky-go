package heap

// heap implements container/heap.Interface.
// It is a private type used by Container.
type heap[T any] struct {
	items []T
	less  func(a, b T) bool
}

// Len returns the number of elements in the heap.
func (instance heap[T]) Len() int {
	return len(instance.items)
}

// Less compares two elements based on the less function.
func (instance heap[T]) Less(indexItemA, indexItemB int) bool {
	return instance.less(instance.items[indexItemA], instance.items[indexItemB])
}

// Swap swaps two elements in the heap.
func (instance heap[T]) Swap(indexItemA, indexItemB int) {
	instance.items[indexItemA], instance.items[indexItemB] = instance.items[indexItemB], instance.items[indexItemA]
}

// Push adds an element to the heap.
// This is used by container/heap.
func (instance *heap[T]) Push(item any) {
	instance.items = append(instance.items, item.(T))
}

// Pop removes and returns the last element from the heap.
// This is used by container/heap.
func (instance *heap[T]) Pop() any {
	lastIndex := len(instance.items) - 1
	deletedItem := instance.items[lastIndex]
	var zero T
	instance.items[lastIndex] = zero
	instance.items = instance.items[:lastIndex]
	return deletedItem
}
