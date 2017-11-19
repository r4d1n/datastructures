package datastructures_test

import (
	"testing"

	"github.com/r4d1n/datastructures"
)

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
