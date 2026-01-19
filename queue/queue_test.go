package queue

import "testing"

func TestQueue(t *testing.T) {
	q := New[string]()

	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}

	q.Enqueue("first")
	q.Enqueue("second")

	if q.Len() != 2 {
		t.Errorf("Expected length 2, got %d", q.Len())
	}

	val, err := q.Peek()
	if err != nil || val != "first" {
		t.Errorf("Expected peek 'first', got '%s', err: %v", val, err)
	}

	val, err = q.Dequeue()
	if err != nil || val != "first" {
		t.Errorf("Expected dequeue 'first', got '%s', err: %v", val, err)
	}

	val, err = q.Dequeue()
	if err != nil || val != "second" {
		t.Errorf("Expected dequeue 'second', got '%s', err: %v", val, err)
	}

	if !q.IsEmpty() {
		t.Error("Queue should be empty")
	}

	_, err = q.Dequeue()
	if err == nil {
		t.Error("Expected error when dequeuing from empty queue")
	}
}

func TestQueue_All_IteratesFIFO(t *testing.T) {
	q := New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	var got []int
	for v := range q.All() {
		got = append(got, v)
	}

	// Should iterate from first (1) to last (3)
	if len(got) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(got))
	}
	if got[0] != 1 || got[1] != 2 || got[2] != 3 {
		t.Errorf("Expected [1 2 3], got %v", got)
	}
}
