// utilization of cpu cache by different data structures

// go test -run none -bench . -benchtime 3s -benchmem

// pkg: Labs/lab_benchmarks/ch_2
// cpu: Intel(R) Core(TM) i5-5250U CPU @ 1.60GHz
// BenchmarkLinkListTraverse-4           97          41356009 ns/op		链表遍历
// BenchmarkColumnTraverse-4             15         270230063 ns/op		列遍历
// BenchmarkRowTraverse-4               169          22363810 ns/op		行遍历
package caching

import "testing"

var fa int

// Capture the time it takes to perform a link list traversal.
func BenchmarkLinkListTraverse(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = LinkedListTraverse()
	}

	fa = a
}

// Capture the time it takes to perform a column traversal.
func BenchmarkColumnTraverse(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = ColumnTraverse()
	}

	fa = a
}

// Capture the time it takes to perform a row traversal.
func BenchmarkRowTraverse(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = RowTraverse()
	}

	fa = a
}
