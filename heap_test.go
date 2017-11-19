package datastructures_test

import (
	"testing"

	"github.com/r4d1n/datastructures"
)

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
