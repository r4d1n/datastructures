package datastructures

// ListNode is the building block for a linked list
type ListNode struct {
	data int
	next *ListNode
}

// NewLinkedList creates a linked list with one node
func NewLinkedList(d int) *ListNode {
	return &ListNode{
		data: d,
		next: nil,
	}
}

// Next returns the next node
func (n *ListNode) Next() *ListNode {
	return n.next
}

// Data returns a node's data
func (n *ListNode) Data() int {
	return n.data
}

// Append a new node to the end of a linked list
func (n *ListNode) Append(d int) *ListNode {
	current := n
	for current.next != nil {
		current = current.next
	}
	current.next = &ListNode{
		data: d,
		next: nil,
	}
	return current.next
}

// Delete a node from a linked list
func (n *ListNode) Delete(d int) *ListNode {
	current := n
	if current.data == d {
		return n.next
	}
	for current.next != nil {
		if current.next.data == d {
			current.next = current.next.next
			return n
		}
		current = current.next
	}
	return n
}

// GetKthFromEnd returns the node k places from the end of the list
func (n *ListNode) GetKthFromEnd(k int) *ListNode {
	behind := n
	ahead := n
	// move the second pointer k places ahead of the first
	for i := 0; i < k-1; i++ {
		ahead = ahead.next
		if ahead == nil {
			return nil
		}
	}

	// move both until the second pointer reaches the end
	// then the first pointer will be k places from the end
	for ahead.next != nil {
		behind = behind.next
		ahead = ahead.next
	}
	return behind
}

// Length returns the distance from a root node to the end of a linked list
func (n *ListNode) Length() int {
	current := n
	counter := 0
	for current != nil {
		counter++
		current = current.next
	}
	return counter
}

// Stack is a Last-In-First-Out data structure
type Stack struct {
	top *ListNode
}

// NewStack returns a new Stack
func NewStack() *Stack {
	var n *ListNode
	return &Stack{top: n}
}

// Pop removes the top item from the Stack and returns it
func (s *Stack) Pop() *ListNode {
	if s.top != nil {
		result := s.top
		s.top = s.top.next
		return result
	}
	return nil
}

// Push adds a new item to the top of the Stack
func (s *Stack) Push(d int) *ListNode {
	n := &ListNode{
		data: d,
		next: s.top,
	}
	s.top = n
	return s.top
}

// Peek returns the top item without removing it from the Stack
func (s *Stack) Peek() int {
	return s.top.data
}

// Length returns the number of items in the Stack
func (s *Stack) Length() int {
	if s.top != nil {
		return s.top.Length()
	}
	return 0
}
