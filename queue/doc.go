// Package queue provides a generic FIFO (First In, First Out) queue implementation.
//
// The queue uses a dynamic slice-based implementation with automatic memory
// optimization to prevent memory leaks in long-running applications.
//
// Example usage:
//
//	q := queue.New[int]()
//	q.Enqueue(10)
//	q.Enqueue(20)
//	q.Enqueue(30)
//
//	val, err := q.Peek()    // Returns 10, nil (without removing)
//	val, err = q.Dequeue()  // Returns 10, nil
//	val, err = q.Dequeue()  // Returns 20, nil
//
//	if q.IsEmpty() {
//	    fmt.Println("Queue is empty")
//	}
//
//	// Iterate values without allocating
//	for v := range q.All() {
//	    fmt.Println(v)
//	}
//
//	// Clear the queue
//	q.Clear()
package queue
