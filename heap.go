package datastructures

import "fmt"

// Heap guarantees access to the minimum or maximum item -- just min for now
type Heap struct {
	data []int
}

// NewHeap instantiates an empty Heap
func NewHeap() *Heap {
	return &Heap{}
}

func getLeftChildIndex(p int) int {
	return 2*p + 1
}

func getRightChildIndex(p int) int {
	return 2*p + 2
}

func getParentIndex(c int) int {
	return (c - 1) / 2
}

// swap items at two given indices
func (h *Heap) swap(i int, j int) error {
	temp := h.data[i]
	h.data[i] = h.data[j]
	h.data[j] = temp
	return nil
}

// Peek returns the top item without removing it
func (h *Heap) Peek() (int, error) {
	if len(h.data) < 1 {
		return 0, fmt.Errorf("no heap data found")
	}
	return h.data[0], nil
}

// Poll removes the top element and returns it
func (h *Heap) Poll() (int, error) {
	if len(h.data) < 1 {
		return 0, fmt.Errorf("no heap data found")
	}
	result := h.data[0]
	h.data[0] = h.data[len(h.data)-1]  // copy last item to the top of the heap
	h.data = h.data[0 : len(h.data)-1] // shrink the slice so last item isn't duplicated
	h.heapifyDown()
	return result, nil
}

// Add inserts a new item into the heap
func (h *Heap) Add(m int) error {
	h.data = append(h.data, m) // add onto the end
	h.heapifyUp()
	return nil
}

func (h *Heap) heapifyDown() error {
	i := 0
	c := getLeftChildIndex(i) // start with the left child
	r := getRightChildIndex(i)
	// check if the right child is actually larger than the left (if both are in the slice)
	if c < len(h.data) && r < len(h.data) && h.data[c] > h.data[r] {
		c = r
	}
	for i < len(h.data) && c < len(h.data) && h.data[i] > h.data[c] {
		h.swap(i, c)
		i++
		c = getLeftChildIndex(i)
		r = getRightChildIndex(i)
		if c < len(h.data) && r < len(h.data) && h.data[c] > h.data[r] {
			c = r
		}
	}
	return nil
}

func (h *Heap) heapifyUp() error {
	i := len(h.data) - 1 // item just added to the end
	p := getParentIndex(i)
	for p < len(h.data) && h.data[p] > h.data[i] {
		h.swap(p, i)
		i = getParentIndex(i)
		p = getParentIndex(i)
	}
	return nil
}

// Length returns the number of items stored in the Heap
func (h *Heap) Length() int {
	return len(h.data)
}
