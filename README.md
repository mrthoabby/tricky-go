# tricky-go

High-performance generic data structures for Go.

## Table of Contents

- [Installation](#installation)
- [Data Structures](#data-structures)
  - [Heap](#heap)
  - [Set](#set)
  - [HashMap](#hashmap)
  - [LinkedList](#linkedlist)
  - [List](#list)
  - [Queue](#queue)
  - [Stack](#stack)
  - [Trie](#trie)
- [Thread Safety](#thread-safety)
- [Documentation](#documentation)
- [License](#license)

---

## Installation

```bash
go get github.com/mrthoabby/tricky-go/...
```

Or install specific packages:

```bash
go get github.com/mrthoabby/tricky-go/heap
go get github.com/mrthoabby/tricky-go/set
go get github.com/mrthoabby/tricky-go/hashmap
go get github.com/mrthoabby/tricky-go/linkedlist
go get github.com/mrthoabby/tricky-go/list
go get github.com/mrthoabby/tricky-go/queue
go get github.com/mrthoabby/tricky-go/stack
go get github.com/mrthoabby/tricky-go/trie
```

---

## Data Structures

### Heap

Priority queue. Smallest or largest item first.

**Use when:** You need items in priority order.

```go
package main

import "github.com/mrthoabby/tricky-go/heap"

type Task struct {
    Name     string
    Priority int
}

func main() {
    // Min-heap: lower priority value first
    h := heap.New(func(a, b Task) bool {
        return a.Priority < b.Priority
    })
    
    h.Push(Task{Name: "Low", Priority: 3})
    h.Push(Task{Name: "High", Priority: 1})
    
    task, _ := h.Pop()  // Gets "High" (priority 1)
}
```

**Methods:**
- `New(less func(a, b T) bool)` - Create heap with comparator (must not be nil)
- `Push(item)` - Add item
- `Pop()` - Get and remove top item
- `Peek()` - Get top item without removing
- `Len()` - Count items
- `Clear()` - Remove all items (does not re-heapify)

[Back to top](#table-of-contents)

---

### Set

Collection of unique items. No duplicates.

**Use when:** You need unique items only.

```go
package main

import "github.com/mrthoabby/tricky-go/set"

func main() {
    s := set.New[string]()
    
    s.Add("apple")
    s.Add("banana")
    s.Add("apple")  // Ignored, already exists
    
    s.Contains("apple")  // true
    s.Remove("banana")
    s.Length()  // 1
}
```

**Methods:**
- `New()` - Create empty set
- `Add(item)` - Add item
- `Remove(item)` - Remove item
- `Contains(item)` - Check if exists
- `Union(other)` - Combine two sets
- `Intersection(other)` - Common items only
- `Length()` - Count items
- `ToSlice()` - Convert to slice
- `All()` - Iterate items without allocation (order not guaranteed)

[Back to top](#table-of-contents)

---

### HashMap

Key/value store with O(1) average lookups and updates.

**Use when:** You need fast key-based access with minimal overhead.

```go
package main

import "github.com/mrthoabby/tricky-go/hashmap"

func main() {
    m := hashmap.New[string, int]()

    m.Set("apple", 3)
    m.Set("banana", 2)

    value, ok := m.Get("apple") // value = 3, ok = true
    _ = value
}
```

**Methods:**
- `New()` - Create empty map
- `NewWithCapacity(capacity)` - Create empty map with reserved capacity
- `Set(key, value)` - Store value for key
- `Get(key)` - Read value and exists flag
- `Delete(key)` - Remove key
- `Contains(key)` - Check if key exists
- `Len()` - Count items
- `Clear()` - Remove all items
- `Keys()` - Get slice of keys (order not guaranteed)
- `Values()` - Get slice of values (order not guaranteed)
- `AsMap()` - Copy to built-in map
- `All()` - Iterate key/value pairs without allocation (order not guaranteed)

[Back to top](#table-of-contents)

---

### LinkedList

List where items point to each other. Go forward or backward.

**Use when:** You need to add/remove items anywhere quickly.

```go
package main

import "github.com/mrthoabby/tricky-go/linkedlist"

func main() {
    list := linkedlist.New[int]()
    
    list.PushBack(1)   // Add to end
    list.PushFront(2)  // Add to start
    
    // Go forward
    for node := list.Front(); node != nil; node = node.Next {
        fmt.Println(node.Value)
    }
    
    // Remove a node
    list.Remove(list.Front())
}
```

**Methods:**
- `New()` - Create empty list
- `PushBack(item)` - Add to end
- `PushFront(item)` - Add to start
- `Front()` - Get first node
- `Back()` - Get last node
- `Remove(node)` - Remove node
- `Len()` - Count items
- `Clear()` - Remove all items
- `ToSlice()` - Convert to slice
- `All()` - Iterate items without allocation

[Back to top](#table-of-contents)

---

### List

Slice-backed list with fast appends and sorting.

**Use when:** You want an array-like list with sortable items.

```go
package main

import "github.com/mrthoabby/tricky-go/list"

func main() {
    l, _ := list.NewWithCapacity[int](3)

    l.Add(3)
    l.Add(1)
    l.Add(2)

    l.Sort(func(a, b int) bool { return a < b })
    l.Remove(1)
}
```

**Methods:**
- `New()` - Create empty list
- `NewWithCapacity(capacity)` - Create empty list with reserved capacity
- `Add(item)` - Add item
- `Remove(index)` - Remove by index
- `Sort(less)` - Sort items
- `Get(index)` - Read by index
- `Len()` - Count items
- `Clear()` - Remove all items
- `ToSlice()` - Copy to slice
- `All()` - Iterate items without allocation

[Back to top](#table-of-contents)

---

### Queue

First in, first out. Like a line at a store.

**Use when:** Process items in order they arrived.

```go
package main

import "github.com/mrthoabby/tricky-go/queue"

func main() {
    q := queue.New[int]()
    
    q.Enqueue(10)  // Add to end
    q.Enqueue(20)
    q.Enqueue(30)
    
    q.Dequeue()  // Gets 10 (first added)
    q.Dequeue()  // Gets 20
    q.Peek()     // See 30 without removing
}
```

**Methods:**
- `New()` - Create empty queue
- `Enqueue(item)` - Add to end
- `Dequeue()` - Remove from front
- `Peek()` - See front without removing
- `IsEmpty()` - Check if empty
- `Len()` - Count items
- `Clear()` - Remove all items
- `All()` - Iterate items without allocation

[Back to top](#table-of-contents)

---

### Stack

Last in, first out. Like a stack of plates.

**Use when:** You need last item added first.

```go
package main

import "github.com/mrthoabby/tricky-go/stack"

func main() {
    s := stack.New[int]()
    
    s.Push(10)  // Add to top
    s.Push(20)
    s.Push(30)
    
    s.Pop()  // Gets 30 (last added)
    s.Pop()  // Gets 20
    s.Peek() // See 10 without removing
}
```

**Methods:**
- `New()` - Create empty stack
- `Push(item)` - Add to top
- `Pop()` - Remove from top
- `Peek()` - See top without removing
- `IsEmpty()` - Check if empty
- `Len()` - Count items
- `Clear()` - Remove all items
- `All()` - Iterate items without allocation

[Back to top](#table-of-contents)

---

### Trie

Key-value prefix tree. Fast prefix search.

**Use when:** You need prefix search or key lookup.

```go
package main

import "github.com/mrthoabby/tricky-go/trie"

func main() {
    t := trie.New[int]()
    
    t.Insert("hello", 1)
    t.Insert("world", 2)
    t.Insert("help", 3)
    
    v, ok := t.Search("hello") // ok=true, v=1
    _, ok = t.Search("hel")    // ok=false (not a complete key)
    t.StartsWith("hel")        // true - prefix exists
    t.Remove("hello")
}
```

**Methods:**
- `New()` - Create empty trie
- `Insert(key, value)` - Add key-value pair
- `Search(key)` - Get value and existence
- `StartsWith(prefix)` - Check if prefix exists
- `Remove(key)` - Remove key
- `All()` - Iterate key-value pairs (order not guaranteed)

[Back to top](#table-of-contents)

---

## Thread Safety

⚠️ **These data structures are NOT thread-safe.**

For concurrent use, add synchronization:

```go
import (
    "sync"
    "github.com/mrthoabby/tricky-go/queue"
)

type SafeQueue[T any] struct {
    q  *queue.Queue[T]
    mu sync.Mutex
}

func (sq *SafeQueue[T]) Enqueue(item T) {
    sq.mu.Lock()
    defer sq.mu.Unlock()
    sq.q.Enqueue(item)
}
```

[Back to top](#table-of-contents)

---

## Documentation

Full documentation via `go doc`:

```bash
go doc github.com/mrthoabby/tricky-go/heap
go doc github.com/mrthoabby/tricky-go/set
go doc github.com/mrthoabby/tricky-go/linkedlist
go doc github.com/mrthoabby/tricky-go/queue
go doc github.com/mrthoabby/tricky-go/stack
go doc github.com/mrthoabby/tricky-go/trie
```

Or view online at [pkg.go.dev](https://pkg.go.dev/github.com/mrthoabby/tricky-go).

[Back to top](#table-of-contents)

---

## License

MIT

[Back to top](#table-of-contents)
