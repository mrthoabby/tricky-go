// Package heap provides a generic priority queue implementation using a binary heap.
//
// The heap uses a caller-provided comparator function and relies on Go's
// standard container/heap package internally for optimal performance.
//
// Example usage:
//
//	type Task struct {
//	    Name     string
//	    Priority int
//	}
//
//	func main() {
//	    // Create a min-heap by comparing priorities
//	    h := heap.New(func(a, b Task) bool {
//	        return a.Priority < b.Priority
//	    })
//
//	    h.Push(Task{Name: "Low", Priority: 3})
//	    h.Push(Task{Name: "High", Priority: 1})
//
//	    task, err := h.Pop() // Returns Task{Name: "High", Priority: 1}, nil
//	    if err != nil {
//	        // Handle error
//	    }
//	}
package heap
