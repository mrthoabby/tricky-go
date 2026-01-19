// Package trie provides a key-value prefix tree (trie) implementation with full Unicode support.
//
// A trie is an efficient data structure for storing and searching strings,
// particularly useful for prefix-based searches and autocomplete functionality.
//
// Example usage:
//
//	t := trie.New[int]()
//	t.Insert("hello", 1)
//	t.Insert("world", 2)
//	t.Insert("help", 3)
//
//	value, found := t.Search("hello")   // Returns 1, true
//	_, found = t.Search("hel")          // Returns 0, false (not a complete key)
//	hasPrefix := t.StartsWith("hel")    // Returns true
//
//	// Iterate key-value pairs (order is not guaranteed)
//	for key, val := range t.All() {
//	    fmt.Println(key, val)
//	}
//
//	// Remove a key
//	removed := t.Remove("hello") // Returns true
package trie
