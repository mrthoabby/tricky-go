// Package list provides a generic, slice-backed list with efficient appends.
//
// A list stores elements in a contiguous slice, making iteration fast and cache
// friendly. Adding items is amortized O(1), removing by index is O(n), and
// sorting uses the provided comparator with O(n log n) complexity.
//
// Example usage:
//
//	items, _ := list.NewWithCapacity[int](3)
//	items.Add(3)
//	items.Add(1)
//	items.Add(2)
//
//	items.Sort(func(a, b int) bool { return a < b })
//	value, _ := items.Remove(1)
//
//	for v := range items.All() {
//	    fmt.Println(v)
//	}
package list
