package bst

import "testing"

var result interface{}

func BenchmarkPut(b *testing.B) {
	bst := New[int]()
	for n := 0; n < b.N; n++ {
		for v, k := range data {
			bst.Put(k, v)
		}
	}
}

func BenchmarkGet(b *testing.B) {
	bst := New[int]()
	for v, k := range data {
		bst.Put(k, v)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, k := range data {
			result = bst.Get(k)
		}
	}
}

func BenchmarkDeleteMin(b *testing.B) {
	bst := New[int]()
	for v, k := range data {
		bst.Put(k, v)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bst.DeleteMin()
	}
}

func BenchmarkDeleteMax(b *testing.B) {
	bst := New[int]()
	for v, k := range data {
		bst.Put(k, v)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bst.DeleteMax()
	}
}

func BenchmarkDelete(b *testing.B) {
	bst := New[int]()
	for v, k := range data {
		bst.Put(k, v)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, k := range data {
			bst.Delete(k)
		}
	}
}

//BenchmarkRank
//BenchmarkSize
//BenchmarkFloor
//BenchmarkCeiling

func BenchmarkKeys(b *testing.B) {
	bst := New[int]()
	for v, k := range data {
		bst.Put(k, v)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		result = bst.Keys()
	}
}

//BenchmarkMin
//BenchmarkMax

func BenchmarkLen(b *testing.B) {
	bst := New[int]()
	for v, k := range data {
		bst.Put(k, v)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		result = bst.Len()
	}
}

func BenchmarkString(b *testing.B) {
	bst := New[int]()
	for v, k := range data {
		bst.Put(k, v)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		result = bst.String()
	}
}
