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

// Delete a node from a linked list and return it or nil if not found
func (n *ListNode) Delete(d int) *ListNode {
	current := n
	if current.data == d {
		return current.next
	}
	for current.next != nil {
		if current.next.data == d {
			current.next = current.next.next
			return n
		}
		current = current.next
	}
	return nil // not found
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
	return &Stack{top: nil}
}

// Pop removes the top item from the Stack and returns a pointer to its data
func (s *Stack) Pop() *int {
	if s.top != nil {
		data := s.top.data
		s.top = s.top.next
		return &data
	}
	return nil
}

// Push adds a new item to the top of the Stack and returns the new length
func (s *Stack) Push(d int) int {
	n := &ListNode{
		data: d,
		next: s.top,
	}
	s.top = n
	return s.Length()
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

// Queue is a First-In-First-Out data structure
type Queue struct {
	first *ListNode
	last  *ListNode
}

// NewQueue returns a new Queue
func NewQueue() *Queue {
	return &Queue{
		first: nil,
		last:  nil,
	}
}

// Enqueue adds an item to the end of the Queue and returns the new length
func (s *Queue) Enqueue(d int) int {
	n := &ListNode{
		data: d,
		next: nil,
	}
	if s.first == nil {
		s.last = n
		s.first = s.last
	} else {
		s.last.next = n
		s.last = s.last.next
	}
	return s.Length()
}

// Dequeue removes an item from the front of the Queue and returns a pointer to its data
func (s *Queue) Dequeue() *int {
	if s.first != nil {
		data := s.first.data
		s.first = s.first.next
		if s.first == nil {
			s.last = nil
		}
		return &data
	}
	return nil
}

// Length returns the number of items in the Queue
func (s *Queue) Length() int {
	if s.first != nil {
		return s.first.Length()
	}
	return 0
}

// BinaryTreeNode is the building block for a binary tree
type BinaryTreeNode struct {
	data  int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

// NewBinaryTree instantiates a binary tree with one node
func NewBinaryTree(d int) *BinaryTreeNode {
	return &BinaryTreeNode{
		data:  d,
		left:  nil,
		right: nil,
	}
}

// Insert adds a new node according to BST sorting principle
func (b *BinaryTreeNode) Insert(d int) *BinaryTreeNode {
	if b == nil {
		return NewBinaryTree(d)
	}
	if d <= b.data {
		b.left = b.left.Insert(d)
	} else {
		b.right = b.right.Insert(d)
	}
	return b
}

// Lookup checks to see if a node with a given value exists in a BST
func (b *BinaryTreeNode) Lookup(d int) bool {
	if b == nil {
		return false
	}
	if d == b.data {
		return true
	} else if d <= b.data {
		return b.left.Lookup(d)
	} else {
		return b.right.Lookup(d)
	}
}

// Left gets the left pointer from a BinaryTreeNode -- mostly a helper for testing
func (b *BinaryTreeNode) Left() *BinaryTreeNode {
	return b.left
}

// Right gets the left pointer from a BinaryTreeNode -- mostly a helper for testing
func (b *BinaryTreeNode) Right() *BinaryTreeNode {
	return b.right
}

// Data gets the data value from a BinaryTreeNode -- mostly a helper for testing
func (b *BinaryTreeNode) Data() int {
	return b.data
}

// NewBST creates a binary search tree from a slice of ints
func NewBST(s []int) *BinaryTreeNode {
	if len(s) < 1 {
		return nil
	}
	root := &BinaryTreeNode{
		data: s[0],
	}
	for _, n := range s[1:] {
		root.Insert(n)
	}
	return root
}
