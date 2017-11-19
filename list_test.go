package datastructures_test

import (
	"testing"

	"github.com/r4d1n/datastructures"
)

// LinkedList tests
func TestLinkedListAppend(t *testing.T) {
	n := datastructures.NewLinkedList(1)
	n.Append(2)
	// confirm result
	if n.Next() == nil {
		t.Errorf("next node should not be nil")
	}
	if n.Next().Data() != 2 {
		t.Errorf("Unexpected result: %v", n.Next().Data())
	}
	n.Append(5)
	if n.Next().Next().Data() != 5 {
		t.Errorf("Unexpected result: %v", n.Next().Next().Data())
	}
}

func TestLinkedListDelete(t *testing.T) {
	n := datastructures.NewLinkedList(1)
	nums := []int{2, 3, 4}
	for _, k := range nums {
		n.Append(k)
	}
	target := 3
	n.Delete(target)
	current := n
	for current.Next() != nil {
		if current.Data() == target {
			t.Errorf("%v should have been deleted from the list", current.Data())
		}
		current = current.Next()
	}
}

func TestLinkedListGetKthFromEnd(t *testing.T) {
	n := datastructures.NewLinkedList(1)
	nums := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, k := range nums {
		n.Append(k)
	}
	result := n.GetKthFromEnd(3) // 8
	if result.Data() != 8 {
		t.Errorf("Unexpected result: %v", result.Data())
	}
}

func TestLinkedListLength(t *testing.T) {
	n := datastructures.NewLinkedList(1)
	nums := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, k := range nums {
		n.Append(k)
	}
	result := n.Length()
	if result != 10 {
		t.Errorf("Unexpected result: %v", result)
	}
}

// Stack tests
func TestStackPush(t *testing.T) {
	s := datastructures.NewStack()
	if s.Length() != 0 {
		t.Errorf("Unexpected length: %v", s.Length())
	}
	// loop through numbers, Push() and test behavior
	nums := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, n := range nums {
		len := s.Push(n)
		if len != i+1 {
			t.Errorf("Unexpected length: %v", len)
		}
		if s.Peek() != n {
			t.Errorf("Unexpected value: %v", s.Peek())
		}
	}
}

func TestStackPop(t *testing.T) {
	s := datastructures.NewStack()
	// Pop() should return nil if the stack is empty
	if s.Length() != 0 {
		t.Errorf("Unexpected length: %v", s.Length())
	}
	result1 := s.Pop()
	if result1 != nil {
		t.Errorf("Unexpected result: %v", result1)
	}
	nums := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, n := range nums {
		s.Push(n)
	}
	for s.Length() > 0 {
		item := s.Pop()
		if *item != nums[s.Length()] {
			t.Errorf("Unexpected result: %v should be %v", *item, nums[s.Length()])
		}
	}
}

// Queue tests
func TestQueueEnqueue(t *testing.T) {
	q := datastructures.NewQueue()
	if q.Length() != 0 {
		t.Errorf("Unexpected length: %v", q.Length())
	}
	// loop through numbers, Enqueue() and test behavior
	nums := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, n := range nums {
		q.Enqueue(n)
		if q.Length() != i+1 {
			t.Errorf("Unexpected length: %v", q.Length())
		}
	}
}

func TestQueueDequeue(t *testing.T) {
	q := datastructures.NewQueue()
	// Dequeue() should return nil if the queue is empty
	if q.Length() != 0 {
		t.Errorf("Unexpected length: %v", q.Length())
	}
	result1 := q.Dequeue()
	if result1 != nil {
		t.Errorf("Unexpected result: %v", result1)
	}
	nums := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, n := range nums {
		q.Enqueue(n)
	}
	index := 0
	for q.Length() > 0 {
		item := q.Dequeue()
		if *item != nums[index] {
			t.Errorf("Unexpected result: %v should be %v", *item, nums[q.Length()])
		}
		index += 1
	}
}
