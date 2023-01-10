package ch_6

// go test -bench='Reuse$' -benchmem -memprofile memprofile.out -cpuprofile profile.out
// go test -run none -bench='Reuse$' -benchmem
/*
	测试slice复用性能提升
*/

import (
	"testing"
)

func genSliceData() []byte {
	s := make([]byte, 10*1024)
	for i := 0; i < 10*1024; i++ {
		s[i] = 'a'
	}
	return s
}

var tmp = genSliceData()

func UnReuse() {
	sUnReuse := make([]byte, 100*1024)
	sUnReuse = append(sUnReuse[:0], tmp...)
}

// slice 不复用每次重新开辟
func BenchmarkSliceUnReuse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UnReuse()
	}
}

var sReuse = make([]byte, 100*1024)

func Reuse() {
	sReuse = sReuse[:0]
	sReuse = append(sReuse, tmp...)
}

// slice 复用
func BenchmarkSliceReuse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reuse()
	}
}
