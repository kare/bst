package bst

import (
	"fmt"
	"strings"
	"testing"
)

var (
	data    = []string{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	lenData = len(data) - 3
)

func TestPut(t *testing.T) {
	b := New[string]()
	if !b.IsEmpty() {
		t.Error("expected empty bst")
	}
	for i, k := range data {
		b.Put(k, fmt.Sprintf("%d", i))
		assertBST(b, t)
	}
	if b.Len() != lenData {
		t.Errorf("expected len %d, but got %d", len(data), b.Len())
	}
	if b.IsEmpty() {
		t.Error("expected non empty bst")
	}
}

func TestKeys(t *testing.T) {
	expected := []struct {
		key   string
		value int
	}{
		{"A", 8},
		{"C", 4},
		{"E", 12},
		{"H", 5},
		{"L", 11},
		{"M", 9},
		{"P", 10},
		{"R", 3},
		{"S", 0},
		{"X", 7},
	}
	b := New[string]()
	for v, k := range data {
		b.Put(k, fmt.Sprintf("%d", v))
	}
	keys := b.Keys()
	for i, td := range expected {
		if keys[i] != td.key {
			t.Errorf("expected key '%v', but got '%v'", td.key, keys[i])
		}
		if b.Get(keys[i]) != fmt.Sprintf("%d", td.value) {
			t.Errorf("expected value '%v', but got '%v'", td.value, b.Get(keys[i]))
		}
	}
}

func TestDeleteMin(t *testing.T) {
	b := New[string]()
	for v, k := range data {
		b.Put(k, fmt.Sprintf("%d", v))
		assertBST(b, t)
	}
	for !b.IsEmpty() {
		t.Log(b)
		b.DeleteMin()
		assertBST(b, t)
	}
}

func TestDeleteMax(t *testing.T) {
	b := New[string]()
	for v, k := range data {
		b.Put(k, fmt.Sprintf("%d", v))
		assertBST(b, t)
	}
	for !b.IsEmpty() {
		t.Log(b)
		b.DeleteMax()
		assertBST(b, t)
	}
}

func TestDelete(t *testing.T) {
	b := New[string]()
	for v, k := range data {
		b.Put(k, fmt.Sprintf("%d", v))
		assertBST(b, t)
	}
	for _, k := range b.Keys() {
		t.Log(b)
		b.Delete(k)
		assertBST(b, t)
	}
	if !b.IsEmpty() {
		t.Error("expected empty bst after delete")
	}
}

func assertBST[T comparable](b *BST[T], t *testing.T) {
	if !isBST(b) {
		t.Error("not in symmetric order")
	}
	if !isSizeConsistent(b) {
		t.Error("subtree counts not consistent")
	}
	if !isRankConsistent(b) {
		t.Error("ranks not consistent")
	}
}

// does this binary tree satisfy symmetric order?
// Note: this test also ensures that data structure is a binary tree since order is strict
func isBST[T comparable](b *BST[T]) bool {
	return isBSTNode(b, b.root, "", "")
}

// is the tree rooted at x a BST with all keys strictly between min and max
// (if min or max is nil, treat as empty constraint)
func isBSTNode[T comparable](b *BST[T], x *node[T], min, max string) bool {
	if x == nil {
		return true
	}
	if min != "" && strings.Compare(x.key, min) <= 0 {
		return false
	}
	if max != "" && strings.Compare(x.key, max) >= 0 {
		return false
	}
	return isBSTNode(b, x.left, min, x.key) && isBSTNode(b, x.right, x.key, max)
}

// are the size fields correct?
func isSizeConsistent[T comparable](b *BST[T]) bool {
	return isSizeConsistentNode(b, b.root)
}

func isSizeConsistentNode[T comparable](b *BST[T], x *node[T]) bool {
	if x == nil {
		return true
	}
	if x.len != b.sizeNode(x.left)+b.sizeNode(x.right)+1 {
		return false
	}
	return isSizeConsistentNode(b, x.left) && isSizeConsistentNode(b, x.right)
}

// check that ranks are consistent
func isRankConsistent[T comparable](b *BST[T]) bool {
	for i := 0; i < b.Len(); i++ {
		key, err := b.Select(i)
		if err != nil {
			fmt.Printf("select error: %v", err)
			return false
		}
		if i != b.Rank(key) {
			return false
		}
	}
	for _, k := range b.Keys() {
		key, err := b.Select(b.Rank(k))
		if err != nil {
			fmt.Printf("select error: %v", err)
			return false
		}
		if strings.Compare(k, key) != 0 {
			return false
		}
	}
	return true
}
