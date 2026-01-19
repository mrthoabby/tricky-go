// Package stack provides a generic LIFO (Last In, First Out) stack implementation.
//
// The stack uses a slice-based implementation for efficient push and pop operations.
//
// Example usage:
//
//	s := stack.New[string]()
//	s.Push("first")
//	s.Push("second")
//	s.Push("third")
//
//	val, err := s.Peek()  // Returns "third", nil (without removing)
//	val, err = s.Pop()    // Returns "third", nil
//	val, err = s.Pop()    // Returns "second", nil
//
//	if !s.IsEmpty() {
//	    fmt.Printf("Stack has %d elements\n", s.Len())
//	}
//
//	// Iterate values without allocating
//	for v := range s.All() {
//	    fmt.Println(v)
//	}
//
//	// Clear the stack
//	s.Clear()
package stack
