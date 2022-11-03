package ch_2

import (
	"strings"
	"testing"
)

func Split(s, seq string) (result []string) {
	i := strings.Index(s, seq)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(seq):]
		i = strings.Index(s, seq)
	}
	result = append(result, s)
	return result
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}

//并行测试
func BenchmarkSplitParallel(b *testing.B) {
	//  b.SetParallelism(4)//设置测试使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("a:b:c", ":")
		}
	})
}
