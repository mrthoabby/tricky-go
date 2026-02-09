// Package hashmap provides a generic hash map implementation.
//
// Example usage:
//
//	package main
//
//	import "github.com/mrthoabby/tricky-go/hashmap"
//
//	func main() {
//		m := hashmap.New[string, int]()
//		m.Set("apples", 3)
//		m.Set("bananas", 2)
//
//		value, ok := m.Get("apples")
//		if ok {
//			_ = value
//		}
//	}
package hashmap
