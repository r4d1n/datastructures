package datastructures

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
