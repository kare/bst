package bst

import (
	"bytes"
	"fmt"
	"strings"
)

type (
	stringQueue []string
	node        struct {
		key   string
		value interface{}
		left  *node
		right *node
		len   int
	}
	// BST represents an ordered symbol table implemented with a binary search tree.
	// Key type is a string and value type is an interface{}.
	// It supports the usual Put, Get, Contains, Delete, Size, and
	// IsEmpty functions.
	// It also provides ordered functions for finding the minimum, maximum,
	// floor, ceiling.
	// It also provides a Keys function for iterating over all of the keys.
	// A symbol table implements the associative array abstraction:
	// when associating a value with a key that is already in the symbol table,
	// the convention is to replace the old value with the new value.
	// Values cannot be nil. Setting the value associated with a key to nil
	// is equivalent to deleting the key from the symbol table.
	// This implementation uses an (unbalanced) binary search tree.
	// The Put, Contains, Remove, Min, Max, Ceiling, Floor, and Rank
	// operations each take linear time in the worst case, if the tree
	// becomes unbalanced.
	// The Size, and IsEmpty operations take constant time.
	// Construction takes constant time.
	BST struct {
		root *node
	}
)

// New returns an empty binary search tree.
func New() *BST {
	return &BST{}
}

// Put inserts key-value pair to tree. If key exists it will update with new
// value. If value is nil key will be removed.
func (b *BST) Put(key string, value interface{}) {
	if value == nil {
		b.Delete(key)
		return
	}
	b.root = b.put(b.root, key, value)
}

func (b *BST) put(x *node, key string, value interface{}) *node {
	if x == nil {
		return &node{
			key:   key,
			value: value,
			len:   1,
		}
	}
	cmp := strings.Compare(key, x.key)
	if cmp < 0 {
		x.left = b.put(x.left, key, value)
	} else if cmp > 0 {
		x.right = b.put(x.right, key, value)
	} else {
		x.value = value
	}
	x.len = 1 + b.sizeNode(x.left) + b.sizeNode(x.right)
	return x
}

// Get returns the value associated with the given key or nil if no such key
// exists.
func (b BST) Get(key string) interface{} {
	return b.get(b.root, key)
}

func (b BST) get(x *node, key string) interface{} {
	if x == nil {
		return nil
	}
	cmp := strings.Compare(key, x.key)
	if cmp < 0 {
		return b.get(x.left, key)
	} else if cmp > 0 {
		return b.get(x.right, key)
	} else {
		return x.value
	}
}

// DeleteMin deletes the smallest key from the tree. If called on an empty tree
// it will silently return.
func (b *BST) DeleteMin() {
	if b.IsEmpty() {
		return
	}
	b.root = b.deleteMin(b.root)
}

func (b *BST) deleteMin(x *node) *node {
	if x.left == nil {
		return x.right
	}
	x.left = b.deleteMin(x.left)
	x.len = b.sizeNode(x.left) + b.sizeNode(x.right) + 1
	return x
}

// DeleteMax deletes the biggest key from the tree. If called on an empty tree
// it will silently return.
func (b *BST) DeleteMax() {
	if b.IsEmpty() {
		return
	}
	b.root = b.deleteMax(b.root)
}

func (b *BST) deleteMax(x *node) *node {
	if x.right == nil {
		return x.left
	}
	x.right = b.deleteMax(x.right)
	x.len = b.sizeNode(x.left) + b.sizeNode(x.right) + 1
	return x
}

// Delete deletes key from the tree.
func (b *BST) Delete(key string) {
	b.root = b.delete(b.root, key)
}

func (b *BST) delete(x *node, key string) *node {
	if x == nil {
		return nil
	}
	cmp := strings.Compare(key, x.key)
	if cmp < 0 {
		x.left = b.delete(x.left, key)
	} else if cmp > 0 {
		x.right = b.delete(x.right, key)
	} else {
		if x.right == nil {
			return x.left
		}
		if x.left == nil {
			return x.right
		}
		t := x
		x = b.min(t.right)
		x.right = b.deleteMin(t.right)
		x.left = t.left
	}
	x.len = b.sizeNode(x.left) + b.sizeNode(x.right) + 1
	return x
}

// Contains returns true when key exists in the tree and false otherwise.
func (b BST) Contains(key string) bool {
	return b.Get(key) != nil
}

func (b BST) selectKey(k int) string {
	if k < 0 || k >= b.Len() {
		return ""
	}
	x := b.selectNode(b.root, k)
	return x.key
}

func (b BST) selectNode(x *node, k int) *node {
	if x == nil {
		return nil
	}
	t := b.sizeNode(x.left)
	if t > k {
		return b.selectNode(x.left, k)
	} else if t < k {
		return b.selectNode(x.right, k-t-1)
	} else {
		return x
	}
}

// Rank returns the number of keys in the symbol table strictly less than key.
func (b BST) Rank(key string) int {
	return b.rankNode(key, b.root)
}

func (b BST) rankNode(key string, x *node) int {
	if x == nil {
		return 0
	}
	cmp := strings.Compare(key, x.key)
	if cmp < 0 {
		return b.rankNode(key, x.left)
	} else if cmp > 0 {
		return 1 + b.sizeNode(x.left) + b.rankNode(key, x.right)
	} else {
		return b.sizeNode(x.left)
	}
}

// Size returns the number of keys in the symbol table in the given range.
func (b BST) Size(lo, hi string) int {
	if strings.Compare(lo, hi) > 0 {
		return 0
	}
	if b.Contains(hi) {
		return b.Rank(hi) - b.Rank(lo) + 1
	}
	return b.Rank(hi) - b.Rank(lo)
}

func (b BST) sizeNode(x *node) int {
	if x == nil {
		return 0
	}
	return x.len
}

// Floor returns the largest key in the symbol table less than or equal to key.
// When key is not found returns an empty string.
func (b BST) Floor(key string) string {
	x := b.floorNode(b.root, key)
	if x == nil {
		return ""
	}
	return x.key
}

func (b BST) floorNode(x *node, key string) *node {
	if x == nil {
		return nil
	}
	cmp := strings.Compare(key, x.key)
	if cmp == 0 {
		return x
	}
	if cmp < 0 {
		return b.floorNode(x.left, key)
	}
	t := b.floorNode(x.right, key)
	if t != nil {
		return t
	}
	return x
}

// Ceiling returns the smallest key in the symbol table greater than or equal to key.
// When key is not found returns an empty string.
func (b BST) Ceiling(key string) string {
	x := b.ceiling(b.root, key)
	if x == nil {
		return ""
	}
	return x.key
}

func (b BST) ceiling(x *node, key string) *node {
	if x == nil {
		return nil
	}
	cmp := strings.Compare(key, x.key)
	if cmp == 0 {
		return x
	}
	if cmp < 0 {
		t := b.ceiling(x.left, key)
		if t != nil {
			return t
		}
		return x
	}
	return b.ceiling(x.right, key)
}

// Keys returns all the keys in the tree.
func (b BST) Keys() []string {
	return b.keys(b.Min(), b.Max())
}

func (b BST) keys(lo, hi string) []string {
	queue := new(stringQueue)
	b.collect(b.root, queue, lo, hi)
	return queue.stringSlice()
}

func (b BST) collect(x *node, queue *stringQueue, lo, hi string) {
	if x == nil {
		return
	}
	cmplo := strings.Compare(lo, x.key)
	cmphi := strings.Compare(hi, x.key)
	if cmplo < 0 {
		b.collect(x.left, queue, lo, hi)
	}
	if cmplo <= 0 && cmphi >= 0 {
		queue.enqueue(x.key)
	}
	if cmphi > 0 {
		b.collect(x.right, queue, lo, hi)
	}
}

// Min returns the smallest key in the symbol table.
// If called on an empty tree it will silently return.
func (b BST) Min() string {
	if b.IsEmpty() {
		return ""
	}
	return b.min(b.root).key
}

func (b BST) min(x *node) *node {
	if x.left == nil {
		return x
	}
	return b.min(x.left)
}

// Max returns the largest key in the symbol table.
// If called on an empty tree it will silently return.
func (b BST) Max() string {
	if b.IsEmpty() {
		return ""
	}
	return b.max(b.root).key
}

func (b BST) max(x *node) *node {
	if x.right == nil {
		return x
	}
	return b.max(x.right)
}

// Len returns the size of the tree.
func (b BST) Len() int {
	return b.sizeNode(b.root)
}

// IsEmpty returns true when the tree is empty and false otherwise.
func (b BST) IsEmpty() bool {
	return b.Len() == 0
}

// String returns a string representation of the tree.
func (b BST) String() string {
	var buf bytes.Buffer
	buf.WriteString("BST{")
	keys := b.Keys()
	lenKeys := len(keys)
	for i, k := range keys {
		buf.WriteString(k)
		buf.WriteString(fmt.Sprintf("%v", b.Get(k)))
		if i+1 != lenKeys {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("}")
	return buf.String()
}

func (q *stringQueue) enqueue(x string) {
	*q = append(*q, x)
}

func (q stringQueue) stringSlice() []string {
	r := make([]string, 0, len(q))
	return append(r, []string(q)...)
}
