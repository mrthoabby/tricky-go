// Package linkedlist provides a generic doubly linked list implementation.
//
// A doubly linked list allows efficient insertion and removal at both ends
// and provides bidirectional traversal. Each node maintains references to
// both the next and previous nodes.
//
// Example usage:
//
//	list := linkedlist.New[int]()
//	list.PushBack(1)
//	list.PushBack(2)
//	list.PushFront(0)
//
//	// Iterate forward
//	for node := list.Front(); node != nil; node = node.Next {
//	    fmt.Println(node.Value)
//	}
//
//	// Remove a specific node
//	node := list.Front()
//	value, err := list.Remove(node)
//	if err != nil {
//	    // Handle error
//	}
//
//	// Iterate values without allocating
//	for v := range list.All() {
//	    fmt.Println(v)
//	}
//
//	// Clear the list
//	list.Clear()
package linkedlist
