package ch_1

import "testing"

func BenchmarkGet(b *testing.B) {
	b.Log("start")
	var addTests = []struct {
		key      string
		value    string
		expected int
	}{
		{"a", "aa", 1},
		{"b", "bb", 2},
		{"c", "cc", 3},
		{"c", "cc", 3},
		{"c", "cc", 3},
		{"d", "dd", 4},
		{"e", "ee", 5},
		{"f", "ff", 6},
		{"g", "gg", 7},
		{"h", "hh", 8},
		{"i", "ii", 9},
		{"j", "jj", 10},
	}
	for _, v := range addTests {
		Add(v.key, v.value)
	}

	//启动内存统计
	b.ReportAllocs()

	//重新计时
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result []string
		for _, v := range addTests {
			value := Get(v.key)
			if value != v.value {
				b.Errorf("get %s:%s, except %s", v.key, value, v.value)
			}
			result = append(result, value)
		}
	}

}
