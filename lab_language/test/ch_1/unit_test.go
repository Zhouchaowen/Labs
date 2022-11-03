package ch_1

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestAddAndGet(t *testing.T) {
	Add("a", "aa")
	fmt.Println(Get("a"))
}

// go test -v -run=TestAdd
func TestAdd(t *testing.T) {
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

	query := rand.Int()
	for _, v := range addTests {
		Add(v.key, v.value)
		t.Logf("[goroutine:%d] add %s:%s", query, v.key, v.value)
		if len(Cash) != v.expected {
			t.Errorf("add %s:%s len = %d; except %d", v.key, v.value, len(Cash), v.expected)
		}
	}
	Clean()
}
