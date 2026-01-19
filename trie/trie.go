package trie

import "iter"

// Node represents a node in the Trie.
type Node[T any] struct {
	children map[rune]*Node[T]
	isEnd    bool
	value    T
}

// Trie implements a Prefix Tree.
// It supports unicode characters.
type Trie[T any] struct {
	root *Node[T]
}

// New creates a new empty Trie.
func New[T any]() *Trie[T] {
	return &Trie[T]{
		root: &Node[T]{children: make(map[rune]*Node[T])},
	}
}

// Insert inserts a value into the Trie associated with the given key.
func (instance *Trie[T]) Insert(key string, value T) {
	current := instance.root
	for _, char := range key {
		if current.children[char] == nil {
			current.children[char] = &Node[T]{children: make(map[rune]*Node[T])}
		}
		current = current.children[char]
	}
	current.isEnd = true
	current.value = value
}

// Search checks if a key exists in the Trie and returns its value.
// Returns (value, true) if found, (zero_value, false) otherwise.
func (instance *Trie[T]) Search(key string) (T, bool) {
	node := instance.searchNode(key)
	if node != nil && node.isEnd {
		return node.value, true
	}
	var zero T
	return zero, false
}

// StartsWith checks if there is any key in the Trie that starts with the given prefix.
func (instance *Trie[T]) StartsWith(prefix string) bool {
	return instance.searchNode(prefix) != nil
}

func (instance *Trie[T]) searchNode(s string) *Node[T] {
	current := instance.root
	for _, char := range s {
		if current.children[char] == nil {
			return nil
		}
		current = current.children[char]
	}
	return current
}

// All returns an iterator over all key-value pairs in the Trie.
// Iteration order is not guaranteed.
func (instance *Trie[T]) All() iter.Seq2[string, T] {
	return func(yield func(string, T) bool) {
		instance.dfs(instance.root, []rune{}, yield)
	}
}

func (instance *Trie[T]) dfs(node *Node[T], currentKey []rune, yield func(string, T) bool) bool {
	if node == nil {
		return true
	}

	if node.isEnd {
		if !yield(string(currentKey), node.value) {
			return false
		}
	}

	for char, child := range node.children {
		if !instance.dfs(child, append(currentKey, char), yield) {
			return false
		}
	}
	return true
}

// Remove removes a key from the Trie.
// Returns true if the key existed and was removed.
func (instance *Trie[T]) Remove(key string) bool {
	runes := []rune(key)
	_, removed := instance.remove(instance.root, runes, 0)
	return removed
}

func (instance *Trie[T]) remove(current *Node[T], key []rune, index int) (bool, bool) {
	if index == len(key) {
		if !current.isEnd {
			return false, false
		}
		current.isEnd = false
		var zero T
		current.value = zero
		return len(current.children) == 0, true
	}

	char := key[index]
	child, exists := current.children[char]
	if !exists {
		return false, false
	}

	shouldDeleteChild, removed := instance.remove(child, key, index+1)
	if removed {
		if shouldDeleteChild {
			delete(current.children, char)
			return !current.isEnd && len(current.children) == 0, true
		}
		return false, true
	}

	return false, false
}
