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

func TestBSTInsert(t *testing.T) {
	b := datastructures.NewBinaryTree(10)
	b.Insert(5)
	b.Insert(15)
	if b.Right().Data() != 15 {
		t.Errorf("Unexpected result: %v should be %v", b.Right(), 15)
	}
	if b.Left().Data() != 5 {
		t.Errorf("Unexpected result: %v should be %v", b.Left(), 5)
	}
	b.Insert(13)
	if b.Right().Left().Data() != 13 {
		t.Errorf("Unexpected result: %v should be %v", b.Right().Left().Data(), 13)
	}
	b.Insert(25)
	if b.Right().Right().Data() != 25 {
		t.Errorf("Unexpected result: %v should be %v", b.Right().Right().Data(), 25)
	}
	b.Insert(22)
	if b.Right().Right().Left().Data() != 22 {
		t.Errorf("Unexpected result: %v should be %v", b.Right().Right().Left().Data(), 22)
	}
	b.Insert(3)
	if b.Left().Left().Data() != 3 {
		t.Errorf("Unexpected result: %v should be %v", b.Left().Left().Data(), 3)
	}
	b.Insert(7)
	if b.Left().Right().Data() != 7 {
		t.Errorf("Unexpected result: %v should be %v", b.Left().Right().Data(), 7)
	}
}

func TestBSTLookup(t *testing.T) {
	b := datastructures.NewBinaryTree(10)
	nums := []int{3, 6, 9, 12, 15, 18, 21}
	for _, n := range nums {
		b.Insert(n)
	}
	// these lookups should return false
	if b.Lookup(500) != false {
		t.Errorf("Unexpected result: %v should be %v", b.Lookup(500), false)
	}
	if b.Lookup(1) != false {
		t.Errorf("Unexpected result: %v should be %v", b.Lookup(1), false)
	}
	// these lookups should return true
	for _, n := range nums {
		if b.Lookup(n) != true {
			t.Errorf("Unexpected result: %v should be %v", b.Lookup(n), true)
		}
	}
}

func TestNewBST(t *testing.T) {
	nums := []int{3, 2, 6, 9, 12, 15, 18, 21}
	b := datastructures.NewBST(nums)
	for _, n := range nums {
		if b.Lookup(n) != true {
			t.Errorf("Unexpected result: %v should be %v", b.Lookup(n), true)
		}
	}
	if b.Data() != 3 {
		t.Errorf("Unexpected result: %v should be %v", b.Data(), 3)
	}
	if b.Right().Data() != 6 {
		t.Errorf("Unexpected result: %v should be %v", b.Right().Data(), 6)
	}
	if b.Left().Data() != 2 {
		t.Errorf("Unexpected result: %v should be %v", b.Left().Data(), 2)
	}
	if b.Right().Right().Data() != 9 {
		t.Errorf("Unexpected result: %v should be %v", b.Right().Right().Data(), 9)
	}
	if b.Right().Left() != nil {
		t.Errorf("Unexpected result: %v should be %v", b.Right().Left(), nil)
	}
	// try with an empty slice
	empty := datastructures.NewBST([]int{})
	if empty != nil {
		t.Errorf("Unexpected result: %v should be %v", empty, nil)
	}
}

// Heap tests
func TestHeapAdd(t *testing.T) {
	h := datastructures.NewHeap()
	nums := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, n := range nums {
		h.Add(n)
		if h.Length() != i+1 {
			t.Errorf("Unexpected length: %v", h.Length())
		}
	}
}

func TestHeapPeek(t *testing.T) {
	h := datastructures.NewHeap()
	nums := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, n := range nums {
		h.Add(n)
		p, e := h.Peek()
		// for now it is just a min heap so check for lowest value
		// in the future should extend with ability to choose min or max
		if e != nil || p != 2 {
			t.Errorf("Unexpected Min: %v", p)
		}
	}
}

func TestHeapPoll(t *testing.T) {
	h := datastructures.NewHeap()
	nums := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, n := range nums {
		h.Add(n)
	}
	var min int
	var err error
	i := 0
	for h.Length() > 0 {
		min, err = h.Poll()
		if err != nil {
			t.Errorf("Unexpected Error: %v", err)
		}
		if h.Length() != len(nums)-(i+1) {
			t.Errorf("Unexpected Length: %v", h.Length())
		}
		if min != nums[i] {
			t.Errorf("Unexpected Min: %v", min)
		}
		i++
	}
}
