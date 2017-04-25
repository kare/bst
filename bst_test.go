package bst

import (
	"strings"
	"testing"
)

var (
	data    = []string{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	lenData = len(data) - 3
)

func TestPut(t *testing.T) {
	b := New()
	if !b.IsEmpty() {
		t.Error("expected empty bst")
	}
	for v, k := range data {
		b.Put(k, v)
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
	b := New()
	for v, k := range data {
		b.Put(k, v)
	}
	keys := b.Keys()
	for i, td := range expected {
		if keys[i] != td.key {
			t.Errorf("expected key '%v', but got '%v'", td.key, keys[i])
		}
		if b.Get(keys[i]) != td.value {
			t.Errorf("expected value '%v', but got '%v'", td.value, b.Get(keys[i]))
		}
	}
}

func TestLevelOrder(t *testing.T) {
	expected := []struct {
		key   string
		value int
	}{
		{"S", 0},
		{"E", 12},
		{"X", 7},
		{"A", 8},
		{"R", 3},
		{"C", 4},
		{"H", 5},
		{"M", 9},
		{"L", 11},
		{"P", 10},
	}
	b := New()
	for v, k := range data {
		b.Put(k, v)
		assertBST(b, t)
	}
	keys := b.LevelOrder()
	for i, td := range expected {
		if keys[i] != td.key {
			t.Errorf("expected key '%v', but got '%v'", td.key, keys[i])
		}
		if b.Get(keys[i]) != td.value {
			t.Errorf("expected value '%v', but got '%v'", td.value, b.Get(keys[i]))
		}
	}
}

func TestDeleteMin(t *testing.T) {
	b := New()
	for v, k := range data {
		b.Put(k, v)
		assertBST(b, t)
	}
	for !b.IsEmpty() {
		t.Log(b)
		b.DeleteMin()
		assertBST(b, t)
	}
}

func TestDeleteMax(t *testing.T) {
	b := New()
	for v, k := range data {
		b.Put(k, v)
		assertBST(b, t)
	}
	for !b.IsEmpty() {
		t.Log(b)
		b.DeleteMax()
		assertBST(b, t)
	}
}

func TestDelete(t *testing.T) {
	b := New()
	for v, k := range data {
		b.Put(k, v)
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

func assertBST(b *BST, t *testing.T) {
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
func isBST(b *BST) bool {
	return isBSTNode(b, b.root, "", "")
}

// is the tree rooted at x a BST with all keys strictly between min and max
// (if min or max is nil, treat as empty constraint)
func isBSTNode(b *BST, x *node, min, max string) bool {
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
func isSizeConsistent(b *BST) bool {
	return isSizeConsistentNode(b, b.root)
}

func isSizeConsistentNode(b *BST, x *node) bool {
	if x == nil {
		return true
	}
	if x.len != b.sizeNode(x.left)+b.sizeNode(x.right)+1 {
		return false
	}
	return isSizeConsistentNode(b, x.left) && isSizeConsistentNode(b, x.right)
}

// check that ranks are consistent
func isRankConsistent(b *BST) bool {
	for i := 0; i < b.Len(); i++ {
		if i != b.Rank(b.selectKey(i)) {
			return false
		}
	}
	for _, key := range b.Keys() {
		if strings.Compare(key, b.selectKey(b.Rank(key))) != 0 {
			return false
		}
	}
	return true
}
