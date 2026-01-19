// Package set provides a generic set implementation using Go maps.
//
// A set is an unordered collection of unique elements. This implementation
// provides O(1) average-case performance for insertion, deletion, and lookup.
//
// Example usage:
//
//	s := set.New[string]()
//	s.Add("apple")
//	s.Add("banana")
//	s.Add("apple") // Duplicate, ignored
//
//	if s.Contains("apple") {
//	    fmt.Println("Found apple")
//	}
//
//	// Set operations
//	s2 := set.New[string]()
//	s2.Add("banana")
//	s2.Add("cherry")
//
//	union := s.Union(s2)        // {apple, banana, cherry}
//	intersection := s.Intersection(s2) // {banana}
//
//	// Iterate values without allocating (order is not guaranteed)
//	for v := range s.All() {
//	    fmt.Println(v)
//	}
package set
